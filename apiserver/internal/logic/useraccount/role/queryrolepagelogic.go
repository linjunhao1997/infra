package role

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRolePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRolePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRolePageLogic {
	return &QueryRolePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryRolePageLogic) QueryRolePage(req *types.QueryRolePageReq) (resp *types.QueryRolePageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
