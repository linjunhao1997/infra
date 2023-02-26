package webapi

import (
	"context"

	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateWebApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateWebApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWebApiLogic {
	return &UpdateWebApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateWebApiLogic) UpdateWebApi(req *types.UpdateWebApiReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
