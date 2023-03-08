package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateWebApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateWebApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWebApiLogic {
	return &UpdateWebApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateWebApiLogic) UpdateWebApi(in *v1.UpdateWebApiReq) (*v1.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &v1.CommonResp{}, nil
}
