package sysadmin

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysadminListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySysadminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysadminListLogic {
	return &QuerySysadminListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySysadminListLogic) QuerySysadminList(req *types.QuerySysadminListReq) (resp *types.QuerySysadminListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
