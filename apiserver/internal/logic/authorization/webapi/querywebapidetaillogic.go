package webapi

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryWebApiDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryWebApiDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryWebApiDetailLogic {
	return &QueryWebApiDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryWebApiDetailLogic) QueryWebApiDetail(req *types.QueryWebApiDetailReq) (resp *types.QueryWebApiDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
