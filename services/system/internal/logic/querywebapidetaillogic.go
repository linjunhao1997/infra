package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryWebApiDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryWebApiDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryWebApiDetailLogic {
	return &QueryWebApiDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryWebApiDetailLogic) QueryWebApiDetail(in *v1.QueryWebApiDetailReq) (*v1.QueryWebApiDetailResp, error) {
	// todo: add your logic here and delete this line

	return &v1.QueryWebApiDetailResp{}, nil
}
