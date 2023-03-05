package authentication

import (
	"context"
	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/pkg/useraccount"
	"infra/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Authenticator[T any] interface {
	Authenticate(ctx context.Context, info T) (string, error)
}

type InternalAuthenticator struct {
	db *gorm.DB
}

type InternalAuthInfo struct {
	UserType string
	Email    string
	Phone    string
	Uid      string
	pwd      string
}

func NewInternalAuthenticator(db *gorm.DB) *InternalAuthenticator {
	return &InternalAuthenticator{
		db: db,
	}
}

func (authenticator *InternalAuthenticator) Authenticate(ctx context.Context, info InternalAuthInfo) (string, error) {
	filter := dao.Filter{Eq: []dao.Eq{{ColumnName: "pwd", ColumnValue: info.pwd}}}
	if info.Email != "" {
		filter.Eq = append(filter.Eq, dao.Eq{ColumnName: "email", ColumnValue: info.Email})
	} else if info.Phone != "" {
		filter.Eq = append(filter.Eq, dao.Eq{ColumnName: "phone", ColumnValue: info.Phone})
	} else if info.Uid != "" {
		filter.Eq = append(filter.Eq, dao.Eq{ColumnName: "uid", ColumnValue: info.Uid})
	}

	authToken := uuid.NewString()
	cacheRegistry := NewUserSessionCacheRegistry(authToken)
	switch info.UserType {
	case types.UserAccountType.Sysadmin:
		sysadmin, err := dao.NewDAO[useraccount.Sysadmin](authenticator.db).FindOne(ctx, filter)
		if err != nil {
			return "", errorx.ErrSysadminAuthFalied
		}

		ver, err := NewAuthVerCacheRegistry(sysadmin.Gid).Get(ctx)
		if err != nil {
			return "", err
		}

		roles, err := sysadmin.GetRoles()
		if err != nil {
			return "", err
		}

		if err := cacheRegistry.Cache(ctx, UserSession{
			Uid:      sysadmin.Uid,
			Gid:      sysadmin.Gid,
			UserType: info.UserType,
			AuthVer:  ver,
			Roles:    roles.Names(),
		}); err != nil {
			return "", err
		}
		return authToken, err
	default:
		return "", errorx.ErrAuthenticationUnsupported
	}
}
