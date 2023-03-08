package logic

import (
	"context"

	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWebElemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateWebElemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWebElemLogic {
	return &CreateWebElemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateWebElemLogic) CreateWebElem(in *v1.CreateWebElemReq) (*v1.CommonIdDataResp, error) {

	return &v1.CommonIdDataResp{}, nil
}
