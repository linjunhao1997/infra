package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteWebElemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteWebElemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWebElemLogic {
	return &DeleteWebElemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteWebElemLogic) DeleteWebElem(in *v1.DeleteWebElemReq) (*v1.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &v1.CommonResp{}, nil
}
