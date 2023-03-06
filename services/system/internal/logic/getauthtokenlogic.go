package logic

import (
	"context"

	"infra/pkg/errorx"
	"infra/services/system/internal/pkg/authentication"
	"infra/services/system/internal/store"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuthTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthTokenLogic {
	return &GetAuthTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuthTokenLogic) GetAuthToken(in *v1.GetAuthTokenReq) (*v1.GetAuthTokenResp, error) {
	var authToken string
	var err error
	if in.InternalAuthMode != nil {
		authToken, err = authentication.NewInternalAuthenticator(store.GormDB).
			Authenticate(l.ctx, authentication.InternalAuthInfo{
				UserType: in.UserType,
				Email:    in.InternalAuthMode.Email,
				Phone:    nil,
				Uid:      in.InternalAuthMode.Uid,
				Pwd:      in.InternalAuthMode.Pwd,
			})
	} else {
		return nil, errorx.ErrAuthenticationAuthModeUnsupported
	}

	if err != nil {
		return nil, errorx.ErrAuthenticationFalied.WrapErr(err)
	}

	return &v1.GetAuthTokenResp{
		Code: 0,
		Msg:  "success",
		Err:  "",
		Data: &v1.GetAuthTokenData{
			AuthToken: authToken,
		},
	}, nil
}
