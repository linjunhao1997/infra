package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryWebApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryWebApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryWebApiListLogic {
	return &QueryWebApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryWebApiListLogic) QueryWebApiList(in *v1.QueryWebApiListReq) (*v1.QueryWebApiListResp, error) {
	// todo: add your logic here and delete this line

	return &v1.QueryWebApiListResp{}, nil
}
