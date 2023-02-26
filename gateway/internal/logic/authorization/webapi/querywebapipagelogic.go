package webapi

import (
	"context"

	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryWebApiPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryWebApiPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryWebApiPageLogic {
	return &QueryWebApiPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryWebApiPageLogic) QueryWebApiPage(req *types.QueryWebApiPageReq) (resp *types.QueryWebApiPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
