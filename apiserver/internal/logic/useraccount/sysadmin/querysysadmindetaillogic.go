package sysadmin

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysadminDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySysadminDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysadminDetailLogic {
	return &QuerySysadminDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySysadminDetailLogic) QuerySysadminDetail(req *types.QuerySysadminDetailReq) (resp *types.QuerySysadminDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
