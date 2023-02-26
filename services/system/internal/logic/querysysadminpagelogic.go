package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysadminPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySysadminPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysadminPageLogic {
	return &QuerySysadminPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuerySysadminPageLogic) QuerySysadminPage(in *v1.QuerySysadminPageReq) (*v1.QuerySysadminPageResp, error) {
	// todo: add your logic here and delete this line

	return &v1.QuerySysadminPageResp{}, nil
}
