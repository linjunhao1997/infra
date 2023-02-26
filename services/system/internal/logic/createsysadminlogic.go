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
	"infra/utils"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSysadminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSysadminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSysadminLogic {
	return &CreateSysadminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSysadminLogic) CreateSysadmin(in *v1.CreateSysadminReq) (*v1.CommonIdDataResp, error) {
	errx := errorx.ErrSysadminCreateFailed
	if len(in.Name) == 0 {
		return nil, errx.Wrap("用户名必传")
	}
	if len(in.Name) > 10 {
		return nil, errx.Wrap("字符超过10个")
	}
	if len(in.Pwd) == 0 {
		return nil, errx.Wrap("密码必传")
	}
	if len(in.Pwd) > 16 {
		return nil, errx.Wrap("密码超过16个")
	}

	mutex, err := utils.NewRedLockUtil(cache.RedisClient).Lock(l.ctx, in.Name, 3*time.Second)
	if err != nil {
		return nil, errx.WrapErr(err)
	}
	defer func() {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			logx.Errorf("unlock failed:%v", err)
		}
	}()

	adminDao := dao.NewDAO[useraccount.Sysadmin](l.svcCtx.DB)
	count, err := adminDao.Count(l.ctx, dao.Filter{
		Eq: []dao.Eq{
			{ColumnName: "name", ColumnValue: in.Name},
		},
	})
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	if count > 0 {
		return nil, errx.Wrap("用户名已存在")
	}

	data := &useraccount.Sysadmin{
		Id: dao.GetID(),
	}

	if err := copier.Copy(data, in); err != nil {
		return nil, errx.WrapErr(err)
	}

	if _, err := adminDao.Create(l.ctx, data); err != nil {
		return nil, errx.WrapErr(err)
	}

	return &v1.CommonIdDataResp{
		Code: 0,
		Msg:  "创建成功",
		Err:  "",
		Data: &v1.CommonIdData{
			Id: data.Id,
		},
	}, nil
}
