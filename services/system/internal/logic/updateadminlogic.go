package logic

import (
	"context"
	"time"

	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/pkg/useraccount"
	"infra/services/system/internal/store"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"
	"infra/types"
	"infra/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminLogic {
	return &UpdateAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAdminLogic) UpdateAdmin(in *v1.UpdateAdminReq) (*v1.CommonResp, error) {
	errx := errorx.ErrAdminUpdateFailed
	if in.Name != nil && len(*in.Name) == 0 {
		return nil, errx.Wrap("用户名不能为空")
	}
	if in.Name != nil && len(*in.Name) > 10 {
		return nil, errx.Wrap("用户名不能超过10个字符")
	}
	if in.Pwd != nil && len(*in.Pwd) == 0 {
		return nil, errx.Wrap("密码不能为空")
	}
	if in.Pwd != nil && len(*in.Pwd) > 16 {
		return nil, errx.Wrap("密码不能超过16个字符")
	}
	adminDao := dao.NewDAO[useraccount.Admin](l.svcCtx.DB)
	original, err := adminDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	latest, err := utils.DeepCopy(original, in)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	mutex, err := utils.NewRedLockUtil(store.RedisClient).Lock(l.ctx, latest.Name, 3*time.Second)
	if err != nil {
		return nil, errx.WrapErr(err)
	}
	defer func() {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			logx.Errorf("unlock failed:%v", err)
		}
	}()

	if original.Name != latest.Name {
		count, err := adminDao.Count(l.ctx, dao.Filter{
			Eq: []dao.Eq{
				{ColumnName: "name", ColumnValue: latest.Name},
				{ColumnName: dao.ColumnDeleted, ColumnValue: types.DeletedType.NotDeleted},
			},
		})
		if err != nil {
			return nil, errx.WrapErr(err)
		}
		if count > 0 {
			return nil, errx.Wrap("用户名已存在")
		}
	}

	if _, err := adminDao.UpdateByModel(l.ctx, latest.Id, latest); err != nil {
		return nil, errx.WrapErr(err)
	}

	return &v1.CommonResp{
		Code: 0,
		Msg:  "修改成功",
		Err:  "",
		Data: nil,
	}, nil
}
