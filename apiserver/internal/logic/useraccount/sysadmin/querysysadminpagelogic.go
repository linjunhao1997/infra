package sysadmin

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysadminPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySysadminPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysadminPageLogic {
	return &QuerySysadminPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySysadminPageLogic) QuerySysadminPage(req *types.QuerySysadminPageReq) (resp *types.QuerySysadminPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
