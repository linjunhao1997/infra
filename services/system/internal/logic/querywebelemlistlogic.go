package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryWebElemListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryWebElemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryWebElemListLogic {
	return &QueryWebElemListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryWebElemListLogic) QueryWebElemList(in *v1.QueryWebElemListReq) (*v1.QueryWebElemListResp, error) {
	// todo: add your logic here and delete this line

	return &v1.QueryWebElemListResp{}, nil
}
