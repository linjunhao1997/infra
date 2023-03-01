package authentication

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"
	v1 "infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthTokenLogic {
	return &GetAuthTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthTokenLogic) GetAuthToken(req *types.GetAuthTokenReq) (resp *types.GetAuthTokenResp, err error) {
	res, err := l.svcCtx.SystemSvc.GetAuthToken(l.ctx, &v1.GetAuthTokenReq{
		InternalAuthMode: &v1.InternalAuthMode{
			Uid:   req.InternalAuthMode.Uid,
			Email: req.InternalAuthMode.Email,
			Pwd:   req.InternalAuthMode.Pwd,
		},
	})
	if err != nil {
		return nil, err
	}

	return &types.GetAuthTokenResp{
		Code: res.Code,
		Msg:  "success",
		Err:  res.Err,
		Data: types.GetAuthTokenData{
			AuthToken: res.Data.AuthToken,
		},
	}, nil
}
