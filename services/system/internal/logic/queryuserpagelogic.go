package logic

import (
	"context"

	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserPageLogic {
	return &QueryUserPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserPageLogic) QueryUserPage(in *v1.QueryUserPageReq) (*v1.QueryUserPageResp, error) {
	// todo: add your logic here and delete this line

	return &v1.QueryUserPageResp{}, nil
}
