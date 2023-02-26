package admin

import (
	"context"

	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
	v1 "infra/services/system/pb/v1"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAdminDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAdminDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAdminDetailLogic {
	return &QueryAdminDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAdminDetailLogic) QueryAdminDetail(req *types.QueryAdminDetailReq) (resp *types.QueryAdminDetailResp, err error) {
	res, err := l.svcCtx.SystemSvc.QueryAdminDetail(l.ctx, &v1.QueryAdminDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	data := &types.AdminData{}
	if err := copier.Copy(data, res.Data); err != nil {
		return nil, err
	}

	return &types.QueryAdminDetailResp{
		Code: res.Code,
		Msg:  res.Msg,
		Err:  res.Err,
		Data: data,
	}, nil
}
