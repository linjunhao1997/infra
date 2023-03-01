package webapi

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryWebApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryWebApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryWebApiListLogic {
	return &QueryWebApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryWebApiListLogic) QueryWebApiList(req *types.QueryWebApiListReq) (resp *types.QueryWebApiListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
