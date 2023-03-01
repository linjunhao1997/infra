package sysadmin

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysadminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysadminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysadminLogic {
	return &DeleteSysadminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysadminLogic) DeleteSysadmin(req *types.DeleteSysadminReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
