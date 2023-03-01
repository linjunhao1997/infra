package admin

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAdminPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAdminPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAdminPageLogic {
	return &QueryAdminPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAdminPageLogic) QueryAdminPage(req *types.QueryAdminPageReq) (resp *types.QueryAdminPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
