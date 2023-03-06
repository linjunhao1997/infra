package authentication

import (
	"context"

	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"
	v1 "infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthTokenByInternalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthTokenByInternalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthTokenByInternalLogic {
	return &GetAuthTokenByInternalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthTokenByInternalLogic) GetAuthTokenByInternal(req *types.InternalAuthReq) (resp *types.GetAuthTokenResp, err error) {
	res, err := l.svcCtx.SystemSvc.GetAuthToken(l.ctx, &v1.GetAuthTokenReq{
		UserType: req.UserType,
		InternalAuthMode: &v1.InternalAuthMode{
			Uid:   req.Uid,
			Email: req.Email,
			Pwd:   req.Pwd,
		},
	})
	if err != nil {
		return nil, err
	}

	return &types.GetAuthTokenResp{
		Code: res.Code,
		Msg:  res.Msg,
		Err:  res.Err,
		Data: types.GetAuthTokenData{
			AuthToken: res.Data.AuthToken,
		},
	}, nil
}
