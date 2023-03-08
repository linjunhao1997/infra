package logic

import (
	"context"

	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/pkg/authorization"
	"infra/services/system/internal/store"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWebApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateWebApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWebApiLogic {
	return &CreateWebApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateWebApiLogic) CreateWebApi(in *v1.CreateWebApiReq) (*v1.CommonIdDataResp, error) {
	id := dao.GetID()
	if _, err := dao.NewDAO[authorization.WebApi](store.GormDB).Create(l.ctx, &authorization.WebApi{
		Id:          id,
		Code:        in.Code,
		Url:         in.Url,
		Disabled:    in.Disabled,
		Description: in.Description,
	}); err != nil {
		return nil, errorx.ErrWebApiCreateFailed.WrapErr(err)
	}

	return &v1.CommonIdDataResp{
		Msg: "success",
		Data: &v1.CommonIdData{
			Id: id,
		},
	}, nil
}
