package webapi

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWebApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateWebApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWebApiLogic {
	return &CreateWebApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWebApiLogic) CreateWebApi(req *types.CreateWebApiReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
