package role

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleListLogic {
	return &QueryRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryRoleListLogic) QueryRoleList(req *types.QueryRoleListReq) (resp *types.QueryRoleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
