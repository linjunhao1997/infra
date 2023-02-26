package sysadmin

import (
	"context"

	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSysadminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSysadminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSysadminLogic {
	return &CreateSysadminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSysadminLogic) CreateSysadmin(req *types.CreateSysadminReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
