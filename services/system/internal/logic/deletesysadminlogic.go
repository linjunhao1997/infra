package logic

import (
	"context"

	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"
	"infra/services/system/pkg/useraccount"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysadminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSysadminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysadminLogic {
	return &DeleteSysadminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSysadminLogic) DeleteSysadmin(in *v1.DeleteSysadminReq) (*v1.CommonResp, error) {
	errx := errorx.ErrSysadminDeleteFailed
	sysadminDao := dao.NewDAO[useraccount.Sysadmin](l.svcCtx.DB)
	role, err := sysadminDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	_, err = sysadminDao.DeleteByIds(l.ctx, role.Id)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	return &v1.CommonResp{
		Code: 0,
		Msg:  "删除成功",
		Err:  "",
		Data: nil,
	}, nil
}
