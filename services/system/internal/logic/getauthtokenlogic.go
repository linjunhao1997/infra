package logic

import (
	"context"
	"time"

	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/pkg/authentication"
	"infra/services/system/internal/pkg/useraccount"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"
	"infra/types"

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
	if in.UserType == types.UserAccountTypeAdmin {
		if in.InternalAuthMode != nil {
			admin, err := dao.NewDAO[useraccount.Admin](l.svcCtx.DB).FindOne(l.ctx, dao.Filter{
				Eq: []dao.Eq{
					{ColumnName: "uid", ColumnValue: *(in.InternalAuthMode.Uid)},
					{ColumnName: "pwd", ColumnValue: in.InternalAuthMode.Pwd},
				},
			})
			if err != nil {
				return nil, errorx.ErrAuthenticationFalied.WrapErr(err)
			}

			ver, err := authentication.GetAuthVer(l.ctx, admin.Gid)
			if err != nil {
				return nil, errorx.ErrAuthenticationFalied.WrapErr(err)
			}

			key := authentication.BuildUserSessionKey(time.Now().String())
			userSession := &authentication.UserSession{
				Uid:     admin.Uid,
				Gid:     admin.Gid,
				AuthVer: ver,
			}
			if err := userSession.Cache(l.ctx, key); err != nil {
				return nil, errorx.ErrAuthenticationFalied.WrapErr(err)
			}

		}
	}

	return &v1.GetAuthTokenResp{}, nil
}
