package user

import (
	"context"

	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserPageLogic {
	return &QueryUserPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserPageLogic) QueryUserPage(req *types.QueryUserPageReq) (resp *types.QueryUserPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
