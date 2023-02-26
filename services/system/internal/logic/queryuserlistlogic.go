package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserListLogic {
	return &QueryUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserListLogic) QueryUserList(in *v1.QueryUserListReq) (*v1.QueryUserListResp, error) {
	// todo: add your logic here and delete this line

	return &v1.QueryUserListResp{}, nil
}
