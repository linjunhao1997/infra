package logic

import (
	"context"

	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAdminLogic {
	return &CreateAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAdminLogic) CreateAdmin(in *v1.CreateAdminReq) (*v1.CommonIdDataResp, error) {
	data, err := l.svcCtx.Service().CreateAdmin(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.CommonIdDataResp{
		Code: 0,
		Msg:  "创建成功",
		Err:  "",
		Data: data,
	}, nil
}
