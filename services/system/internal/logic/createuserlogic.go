package logic

import (
	"context"

	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *v1.CreateUserReq) (*v1.CommonIdDataResp, error) {
	// todo: add your logic here and delete this line

	return &v1.CommonIdDataResp{}, nil
}
