package useraccountcache

/*
import (
	"context"
	"fmt"
	"infra/pkg/dao"
	"infra/server"
	"infra/services/system/pkg/useraccount"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type UserInfo struct {
	Uid   string
	Roles []string
}

//uri: []roles

func buildRolesVerKey(gid string) string {
	return "system::roles::gid::" + gid + "::ver"
}

func buildSessionKey(uid string) string {
	return "system::session::uid::" + uid
}

func GetAuthVer(ctx context.Context, gid string) (string, error) {
	ver, err := server.RedisClient.Get(ctx, buildAuthVerKey(gid)).Result()
	if err != nil {
		return "", err
	}
	return ver, nil
}

func GetRoles(ctx context.Context, gid string) ([]string, error) {
	res := make([]string, 0)
	d := dao.NewDAO[useraccount.Role](server.GormDB)
	var cacheErr error
	ver, cacheErr := GetAuthVersion(ctx, gid)
	if cacheErr == nil {
		res, cacheErr = server.RedisClient.SMembers(ctx, "system::useraccount::"+gid+"::roles::ver::"+ver).Result()
	}

	if cacheErr != nil {
		logx.Errorf("GetRoles err %v", cacheErr)
	}

	if len(res) != 0 {
		return res, nil
	}

	datas, err := d.Find(ctx, dao.Filter{
		Eq: []dao.Eq{{
			ColumnName:  "gid",
			ColumnValue: gid,
		}},
	}, nil)
	if err != nil {
		return nil, err
	}

	for _, e := range datas {
		res = append(res, e.Name)
	}

	threading.GoSafe(func() {
		verInt, err := server.RedisClient.Incr(ctx, "system::useraccount::auth::ver").Result()
		if err != nil {
			logx.Errorf("Incr ver err %v", err)
		}
		if _, err := server.RedisClient.SAdd(ctx, "system::useraccount::"+gid+"::roles::ver::"+ver, res).Result(); err != nil {
			logx.Errorf("Incr ver err %v", err)
		}
		if _, err := server.RedisClient.Expire(ctx, fmt.Sprintf("system::useraccount::%s::roles::ver::%d", gid, verInt), time.Minute*30).Result(); err != nil {
			logx.Errorf("Incr ver err %v", err)
		}
	})

	return res, nil
}

func GetSysAdminInfo(ctx context.Context, uid string) (*useraccount.Sysadmin, error) {
	res := &useraccount.Sysadmin{}
	ver, err := GetAuthVer(ctx, uid)
	if err == nil {
		logx.Errorf("getauthver err %v", err)
		server.RedisClient.Get(ctx, buildSessionKey(ver, uid)).Scan(res)
	}

	dao.NewDAO[useraccount.Sysadmin](server.GormDB).Native().Find(res)
}

func GetAdminInfo(ctx context.Context, gid, uid string) *useraccount.Admin {

} */
