package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuthTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthTokenLogic {
	return &GetAuthTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuthTokenLogic) GetAuthToken(in *v1.GetAuthTokenReq) (*v1.GetAuthTokenResp, error) {
	// todo: add your logic here and delete this line

	return &v1.GetAuthTokenResp{}, nil
}
