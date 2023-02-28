package svc

import (
	"context"
	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/pkg/useraccount"
	"infra/services/system/internal/store"
	v1 "infra/services/system/pb/v1"
	"infra/utils"
	"time"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SystemService interface {
	CreateAdmin(ctx context.Context, req *v1.CreateAdminReq) (*v1.CommonIdData, error)
}

type SystemServiceImpl struct {
	ctx *ServiceContext
}

func NewSystemService(ctx *ServiceContext) SystemService {
	return &SystemServiceImpl{
		ctx: ctx,
	}
}

func (impl SystemServiceImpl) CreateAdmin(ctx context.Context, req *v1.CreateAdminReq) (*v1.CommonIdData, error) {
	errx := errorx.ErrAdminCreateFailed
	if len(req.Name) == 0 {
		return nil, errx.Wrap("用户名必传")
	}
	if len(req.Name) > 10 {
		return nil, errx.Wrap("字符超过10个")
	}
	if len(req.Pwd) == 0 {
		return nil, errx.Wrap("密码必传")
	}
	if len(req.Pwd) > 16 {
		return nil, errx.Wrap("密码超过16个")
	}

	mutex, err := utils.NewRedLockUtil(store.RedisClient).Lock(ctx, req.Name, 3*time.Second)
	if err != nil {
		return nil, errx.WrapErr(err)
	}
	defer func() {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			logx.Errorf("unlock failed:%v", err)
		}
	}()

	adminDao := dao.NewDAO[useraccount.Admin](impl.ctx.DB)
	count, err := adminDao.Count(ctx, dao.Filter{
		Eq: []dao.Eq{
			{ColumnName: "name", ColumnValue: req.Name},
		},
	})
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	if count > 0 {
		return nil, errx.Wrap("用户名已存在")
	}

	data := &useraccount.Admin{
		Id: dao.GetID(),
	}

	if err := copier.Copy(data, req); err != nil {
		return nil, errx.WrapErr(err)
	}

	if _, err := adminDao.Create(ctx, data); err != nil {
		return nil, errx.WrapErr(err)
	}

	return &v1.CommonIdData{
		Id: data.Id,
	}, nil
}
