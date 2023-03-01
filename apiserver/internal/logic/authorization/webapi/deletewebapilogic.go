package webapi

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteWebApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteWebApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWebApiLogic {
	return &DeleteWebApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteWebApiLogic) DeleteWebApi(req *types.DeleteWebApiReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
