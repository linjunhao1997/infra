package sysadmin

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysadminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysadminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysadminLogic {
	return &UpdateSysadminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysadminLogic) UpdateSysadmin(req *types.UpdateSysadminReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
