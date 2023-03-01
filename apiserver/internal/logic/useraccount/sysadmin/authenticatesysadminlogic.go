package sysadmin

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateSysadminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticateSysadminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateSysadminLogic {
	return &AuthenticateSysadminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticateSysadminLogic) AuthenticateSysadmin(req *types.AuthenticateSysadmin) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
