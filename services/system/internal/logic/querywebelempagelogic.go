package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryWebElemPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryWebElemPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryWebElemPageLogic {
	return &QueryWebElemPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryWebElemPageLogic) QueryWebElemPage(in *v1.QueryWebElemPageReq) (*v1.QueryWebElemPageResp, error) {
	// todo: add your logic here and delete this line

	return &v1.QueryWebElemPageResp{}, nil
}
