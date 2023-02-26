package admin

import (
	"context"

	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAdminListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAdminListLogic {
	return &QueryAdminListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAdminListLogic) QueryAdminList(req *types.QueryAdminListReq) (resp *types.QueryAdminListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
