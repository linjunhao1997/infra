package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryWebElemDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryWebElemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryWebElemDetailLogic {
	return &QueryWebElemDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryWebElemDetailLogic) QueryWebElemDetail(in *v1.QueryWebElemDetailReq) (*v1.QueryWebElemDetailResp, error) {
	// todo: add your logic here and delete this line

	return &v1.QueryWebElemDetailResp{}, nil
}
