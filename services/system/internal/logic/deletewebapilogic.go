package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteWebApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteWebApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWebApiLogic {
	return &DeleteWebApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteWebApiLogic) DeleteWebApi(in *v1.DeleteWebApiReq) (*v1.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &v1.CommonResp{}, nil
}
