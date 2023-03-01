package sysmetadata

import (
	"context"
	"errors"
	"fmt"
	"infra/services/system/internal/store"
	"infra/types"
	"strings"
	"time"

	gutils "infra/utils"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	SysMetadataUtil                          = newSysMetadataUtils(store.GormDB, store.RedisClient)
	SysMetadataExpiration      time.Duration = 0
	errSysMetadataCacheEmpty                 = errors.New("sysmetadata cache empty")
	errSysMetadataCacheExpired               = errors.New("sysmetadata cache expired")
)

type SysMetadataUtils struct {
	db          *gorm.DB
	redisClient *redis.Client
	cacheUtils  SysMetadataCacheUtils
}

func newSysMetadataUtils(db *gorm.DB, redisClient *redis.Client) *SysMetadataUtils {
	return &SysMetadataUtils{
		db:          db,
		redisClient: redisClient,
	}
}

func (utils *SysMetadataUtils) GetSysMetadata(ctx context.Context, typeName, code string) (*SysMetadata, error) {
	loadKey := utils.cacheUtils.buildCacheTypeCodeKey(typeName, code)

	// 先查缓存
	data, err := utils.cacheUtils.GetSysMetadataFromCache(ctx, typeName, code)
	if err != nil {
		logx.Infof(`get key %s failed: %v`, loadKey, err)
	} else {
		return data, nil
	}

	// 使用分布式锁防止缓存击穿
	lockKey := utils.cacheUtils.BuildLockKey(typeName, code)
	mutex, err := gutils.NewRedLockUtil(store.RedisClient).Lock(ctx, lockKey, 3*time.Second)
	if err != nil {
		return nil, err
	}
	defer mutex.Unlock()

	// 获得锁之后再查一次缓存
	data, err = utils.cacheUtils.GetSysMetadataFromCache(ctx, typeName, code)
	if err != nil {
		logx.Infof(`get key %s failed: %v`, loadKey, err)
	} else {
		return data, nil
	}

	// 缓存中仍然不存在则从数据库获取数据
	data = &SysMetadata{}
	if err := utils.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "SHARE"}).Debug().
			Take(data,
				fmt.Sprintf(`%s = ? and %s = ? and %s = ? and %s = ?`,
					ColumnSysMetadataType,
					ColumnSysMetadataCode,
					ColumnSysMetadataIsDeleted,
					ColumnSysMetadataIsOnline),
				typeName, code, types.DeletedType.NotDeleted, true).Error; err != nil {
			return err
		}

		threading.GoSafe(func() {
			if err := utils.cacheUtils.RefreshSysMetadata(context.Background(), data); err != nil {
				logx.Errorf("RefreshSysMetadata: %v", err)
			}
		})

		return nil
	}); err != nil {
		return nil, err
	}

	return data, nil
}

// GetSysMetadataList 获取系统元数据列表
func (utils *SysMetadataUtils) GetSysMetadataList(ctx context.Context, typeName string) (SysMetadataList, error) {
	// 从缓存获取
	cacheList, err := utils.cacheUtils.GetSysMetadataListFromCache(ctx, typeName)
	if err != nil {
		logx.Infof("getSysMetadataListFromCache: %v", err)
	} else {
		return cacheList, nil
	}

	// 使用分布式锁防止缓存击穿

	lockKey := utils.cacheUtils.BuildLockKey(typeName, "")
	mutex, err := gutils.NewRedLockUtil(store.RedisClient).Lock(ctx, lockKey, 3*time.Second)
	if err != nil {
		return nil, err
	}
	defer func() {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			logx.Errorf("unlock failed:%v", err)
		}
	}()

	// 尝试再次获取缓存
	cacheList, err = utils.cacheUtils.GetSysMetadataListFromCache(ctx, typeName)
	if err != nil {
		logx.Infof("GetSysMetadataListFromCache: %v", err)
	} else {
		return cacheList, nil
	}

	// 只能从数据库取
	res := SysMetadataList{}
	if err := utils.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "SHARE"}).Debug().
			Order(clause.OrderByColumn{Column: clause.Column{Name: ColumnSysMetadataCode}, Desc: false}).
			Find(&res,
				fmt.Sprintf(`%s = ? and %s = ? and %s = ?`,
					ColumnSysMetadataType,
					ColumnSysMetadataIsDeleted,
					ColumnSysMetadataIsOnline),
				typeName, types.DeletedType.NotDeleted, true).
			Error; err != nil {
			return err
		}

		threading.GoSafe(func() {
			if err := utils.cacheUtils.RefreshSysMetadataList(context.Background(), res); err != nil {
				logx.Errorf("RefreshSysMetadataList: %v", err)
			}
		})

		return nil
	}); err != nil {
		return nil, err
	}

	return res, nil
}

func (utils *SysMetadataUtils) redisPip() redis.Pipeliner {
	return utils.redisClient.Pipeline()
}

func (utils *SysMetadataUtils) GetErrorCodeMsg(code string) string {
	res, err := utils.GetSysMetadata(context.Background(), ERROR_TYPE, code)
	if err != nil || res == nil {
		return ""
	}
	return res.Value
}

// SysMetadataCacheUtils 系统元数据缓存工具类
type SysMetadataCacheUtils struct {
	redisClient *redis.Client
}

// NewSysMetadataCacheUtils 初始化系统元数据缓存工具类
func NewSysMetadataCacheUtils(redisClient *redis.Client) *SysMetadataCacheUtils {
	return &SysMetadataCacheUtils{
		redisClient: redisClient,
	}
}

// removeCacheByType 移除缓存
func (utils *SysMetadataCacheUtils) removeCacheByType(ctx context.Context, pip redis.Pipeliner, typeName string) {
	typeKey := utils.buildCacheTypeKey(typeName)
	utils.removeCacheByKey(ctx, pip, typeKey)
	utils.removeCacheByKeyPattern(ctx, pip, typeKey+":*")
}

// removeCacheByTypeAndCode 移除缓存
func (utils *SysMetadataCacheUtils) removeCacheByTypeAndCode(ctx context.Context, pip redis.Pipeliner, typeName, code string) {
	pip.ZRem(ctx, utils.buildCacheTypeKey(typeName), code)
	utils.removeCacheByKey(ctx, pip, utils.buildCacheTypeCodeKey(typeName, code))
}

// removeCacheByKey 执行redis移除缓存逻辑
func (utils *SysMetadataCacheUtils) removeCacheByKey(ctx context.Context, pip redis.Pipeliner, key string) {
	pip.Del(ctx, key)
}

// removeCacheByKeyPattern 执行redis移除缓存逻辑
func (utils *SysMetadataCacheUtils) removeCacheByKeyPattern(ctx context.Context, pip redis.Pipeliner, pattern string) {
	itr := SysMetadataUtil.redisClient.Scan(ctx, 0, pattern, 1000).Iterator()
	for itr.Next(ctx) {
		pip.Del(ctx, itr.Val())
	}
}

// buildCacheTypeCodeKey 构造key
func (utils *SysMetadataCacheUtils) buildCacheTypeCodeKey(typeName, code string) string {
	return utils.buildKey(SYSMETADATA_PREFIX, typeName, code)
}

// buildCacheTypeKey 构造key
func (utils *SysMetadataCacheUtils) buildCacheTypeKey(typeName string) string {
	return utils.buildKey(SYSMETADATA_PREFIX, typeName, "")
}

// buildKey 构造key
func (utils *SysMetadataCacheUtils) buildKey(prefix, typeName, code string) string {
	names := []string{typeName, code}
	joins := make([]string, 0)
	for _, e := range names {
		if strings.TrimSpace(e) != "" {
			joins = append(joins, e)
		}
	}
	return prefix + strings.Join(joins, ":")
}

// Difference 比对codeSet
func (utils *SysMetadataCacheUtils) Difference(codeSet1, codeSet2 []string) []string {
	set1 := make(map[string]struct{}, 0)
	for _, e := range codeSet1 {
		set1[e] = struct{}{}
	}
	set2 := make(map[string]struct{}, 0)
	for _, e := range codeSet2 {
		set2[e] = struct{}{}
	}

	result := make([]string, 0)
	for k := range set1 {
		if _, ok := set2[k]; !ok {
			result = append(result, k)
		}
	}

	return result
}

// RefreshSysMetadataList 刷新系统元数据列表
func (utils *SysMetadataCacheUtils) RefreshSysMetadataList(ctx context.Context, list SysMetadataList) error {
	pip := SysMetadataUtil.redisPip()
	utils.refreshSysMetadataList(ctx, pip, list)
	_, err := pip.Exec(ctx)
	return err
}

// RemoveCacheByTypeAndCode 移除系统元数据缓存
func (utils *SysMetadataCacheUtils) RemoveCacheByTypeAndCode(ctx context.Context, typeName, code string) error {
	pip := SysMetadataUtil.redisPip()
	utils.removeCacheByTypeAndCode(ctx, pip, typeName, code)
	_, err := pip.Exec(ctx)
	return err
}

// RemoveCacheByType 移除系统元数据缓存
func (utils *SysMetadataCacheUtils) RemoveCacheByType(ctx context.Context, typeName string) error {
	pip := SysMetadataUtil.redisPip()
	utils.removeCacheByType(ctx, pip, typeName)
	_, err := pip.Exec(ctx)
	return err
}

// RefreshAllSysMetadata 刷新所有系统元数据缓存
func (utils *SysMetadataCacheUtils) RefreshAllSysMetadata(ctx context.Context) error {
	// 所有类型
	typeList := SysMetadataList{}
	if err := SysMetadataUtil.db.Order(clause.OrderByColumn{Column: clause.Column{Name: ColumnSysMetadataCode}, Desc: false}).Find(&typeList,
		ColumnSysMetadataType+"=? and "+ColumnSysMetadataIsDeleted+"=?", ROOT_TYPE, 0).Error; err != nil {
		return err
	}

	// 遍历类型
	for _, typeName := range typeList.CodeList() {
		if err := SysMetadataUtil.db.Transaction(func(tx *gorm.DB) error {
			list := SysMetadataList{}
			if err := tx.Clauses(clause.Locking{Strength: "SHARE"}).
				Order(clause.OrderByColumn{Column: clause.Column{Name: ColumnSysMetadataCode}, Desc: false}).
				Find(&list,
					ColumnSysMetadataType+"=? and "+
						ColumnSysMetadataIsDeleted+"=?"+
						ColumnSysMetadataIsOnline+"=?",
					typeName, types.DeletedType.NotDeleted, true).Error; err != nil {
				return err
			}
			return utils.RefreshSysMetadataList(ctx, list)
		}); err != nil {
			return err
		}
	}

	return nil
}

// GetSysMetadataFromCache 从缓存获取数据
func (utils *SysMetadataCacheUtils) GetSysMetadataFromCache(ctx context.Context, typeName, code string) (*SysMetadata, error) {
	data := &SysMetadata{}
	if err := SysMetadataUtil.redisClient.Get(ctx, utils.buildCacheTypeCodeKey(typeName, code)).Scan(data); err != nil {
		return nil, err
	}
	return data, nil
}

// GetSysMetadataListFromCache 从缓存获取数据
func (utils *SysMetadataCacheUtils) GetSysMetadataListFromCache(ctx context.Context, typeName string) (SysMetadataList, error) {
	codeSet, err := utils.GetSysMetadataCodeSetFormCache(ctx, typeName)
	if err != nil {
		return nil, err
	}
	if len(codeSet) == 0 {
		return nil, errSysMetadataCacheEmpty
	}

	cmds := make([]*redis.StringCmd, len(codeSet))
	cmder, _ := SysMetadataUtil.redisClient.Pipelined(ctx, func(pip redis.Pipeliner) error {
		for _, code := range codeSet {
			cmds = append(cmds, pip.Get(ctx, utils.buildCacheTypeCodeKey(typeName, code)))
		}
		return nil
	})

	cacheList := SysMetadataList{}
	for _, cmd := range cmder {
		cmd, ok := cmd.(*redis.StringCmd)
		if !ok {
			continue
		}
		metadata := &SysMetadata{}
		if err := cmd.Scan(metadata); err != nil {
			if errors.Is(err, redis.Nil) {
				continue
			}
			return nil, err
		}
		cacheList = append(cacheList, metadata)
	}

	missCodeSet := utils.Difference(codeSet, cacheList.CodeList())
	if len(missCodeSet) == 0 {
		return cacheList, nil
	}

	return nil, errSysMetadataCacheExpired
}

func (utils *SysMetadataCacheUtils) GetSysMetadataCodeSetFormCache(ctx context.Context, typeName string) ([]string, error) {
	return SysMetadataUtil.redisClient.ZRange(ctx, utils.buildCacheTypeKey(typeName), 0, -1).Result()
}

// RefreshSysMetadata 刷新系统元数据
func (utils *SysMetadataCacheUtils) RefreshSysMetadata(ctx context.Context, data *SysMetadata) error {
	_, err := SysMetadataUtil.redisClient.Set(ctx, utils.buildCacheTypeCodeKey(data.Type, data.Code), data, SysMetadataExpiration).Result()
	return err
}

func (utils *SysMetadataCacheUtils) refreshSysMetadataList(ctx context.Context, pip redis.Pipeliner, list SysMetadataList) {
	for typeName, sysMetadataList := range list.PurgeDeleted().FilterOnline().GroupByType() {
		utils.removeCacheByType(ctx, pip, typeName)
		typeKey := utils.buildCacheTypeKey(typeName)
		for i, e := range sysMetadataList {
			pip.ZAdd(ctx, typeKey, &redis.Z{
				Score:  float64(i),
				Member: e.Code,
			})
			pip.Set(ctx, utils.buildCacheTypeCodeKey(typeName, e.Code), e, SysMetadataExpiration)
		}
	}
}

// BuildLockKey 构造缓存锁的key
func (utils *SysMetadataCacheUtils) BuildLockKey(typeName, code string) string {
	return utils.buildKey(SYSMETADATA_LOCK_PREFIX, typeName, code)
}
