package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateWebElemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateWebElemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWebElemLogic {
	return &UpdateWebElemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateWebElemLogic) UpdateWebElem(in *v1.UpdateWebElemReq) (*v1.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &v1.CommonResp{}, nil
}
