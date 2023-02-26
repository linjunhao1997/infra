package authentication

import (
	"context"

	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthTokenLogic {
	return &GetAuthTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthTokenLogic) GetAuthToken(req *types.GetAuthTokenReq) (resp *types.GetAuthTokenResp, err error) {
	// todo: add your logic here and delete this line

	return
}
