package role

import (
	"context"

	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRoleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleDetailLogic {
	return &QueryRoleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryRoleDetailLogic) QueryRoleDetail(req *types.QueryRoleDetailReq) (resp *types.QueryRoleDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
