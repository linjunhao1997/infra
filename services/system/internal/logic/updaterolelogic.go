package logic

import (
	"context"
	"time"

	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"
	"infra/services/system/pkg/cache"
	"infra/services/system/pkg/useraccount"
	"infra/types"
	"infra/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *v1.UpdateRoleReq) (*v1.CommonResp, error) {
	errx := errorx.ErrRoleUpdateFailed
	if in.Name != nil && len(*in.Name) == 0 {
		return nil, errx.Wrap("角色名不能为空")
	}
	if in.Name != nil && len(*in.Name) > 10 {
		return nil, errx.Wrap("角色名不能超过10个字符")
	}

	roleDao := dao.NewDAO[useraccount.Role](l.svcCtx.DB)
	original, err := roleDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	latest, err := utils.DeepCopy(original, in)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	mutex, err := utils.NewRedLockUtil(cache.RedisClient).Lock(l.ctx, latest.Name, 3*time.Second)
	if err != nil {
		return nil, errx.WrapErr(err)
	}
	defer func() {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			logx.Errorf("unlock failed:%v", err)
		}
	}()

	if original.Name != latest.Name {
		count, err := roleDao.Count(l.ctx, dao.Filter{
			Eq: []dao.Eq{
				{ColumnName: "name", ColumnValue: latest.Name},
				{ColumnName: dao.ColumnDeleted, ColumnValue: types.DeletedType.NotDeleted},
			},
		})
		if err != nil {
			return nil, errx.WrapErr(err)
		}
		if count > 0 {
			return nil, errx.Wrap("角色名已存在")
		}
	}

	if _, err := roleDao.UpdateByModel(l.ctx, latest.Id, latest); err != nil {
		return nil, errx.WrapErr(err)
	}

	return &v1.CommonResp{
		Code: 0,
		Msg:  "修改成功",
		Err:  "",
		Data: nil,
	}, nil
}
