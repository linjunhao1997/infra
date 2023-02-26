package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserDetailLogic {
	return &QueryUserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserDetailLogic) QueryUserDetail(in *v1.QueryUserDetailReq) (*v1.QueryUserDetailResp, error) {
	// todo: add your logic here and delete this line

	return &v1.QueryUserDetailResp{}, nil
}
