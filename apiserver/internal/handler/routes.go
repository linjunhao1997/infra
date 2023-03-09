// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	authentication "infra/apiserver/internal/handler/authentication"
	authorizationwebapi "infra/apiserver/internal/handler/authorization/webapi"
	useraccountadmin "infra/apiserver/internal/handler/useraccount/admin"
	useraccountrole "infra/apiserver/internal/handler/useraccount/role"
	useraccountsysadmin "infra/apiserver/internal/handler/useraccount/sysadmin"
	useraccountuser "infra/apiserver/internal/handler/useraccount/user"
	"infra/apiserver/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/authentication/getauthtoken",
				Handler: authentication.GetAuthTokenHandler(serverCtx),
			},
		},
		rest.WithPrefix("/system/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthorityMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/authorization/webapi/:id",
					Handler: authorizationwebapi.QueryWebApiDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/authorization/webapi/list",
					Handler: authorizationwebapi.QueryWebApiListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/authorization/webapi/page",
					Handler: authorizationwebapi.QueryWebApiPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authorization/webapi",
					Handler: authorizationwebapi.CreateWebApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/authorization/webapi/:id",
					Handler: authorizationwebapi.UpdateWebApiHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/authorization/webapi/:id",
					Handler: authorizationwebapi.DeleteWebApiHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthorityMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/useraccount/admin/:id",
					Handler: useraccountadmin.QueryAdminDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/useraccount/admin/list",
					Handler: useraccountadmin.QueryAdminListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/useraccount/admin/page",
					Handler: useraccountadmin.QueryAdminPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/useraccount/admin",
					Handler: useraccountadmin.CreateAdminHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/useraccount/admin/:id",
					Handler: useraccountadmin.UpdateAdminHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/useraccount/admin/:id",
					Handler: useraccountadmin.DeleteAdminHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthorityMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/useraccount/role/:id",
					Handler: useraccountrole.QueryRoleDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/useraccount/role/list",
					Handler: useraccountrole.QueryRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/useraccount/role/page",
					Handler: useraccountrole.QueryRolePageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/useraccount/role",
					Handler: useraccountrole.CreateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/useraccount/role/:id",
					Handler: useraccountrole.UpdateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/useraccount/role/:id",
					Handler: useraccountrole.DeleteRoleHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthorityMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/useraccount/sysadmin/:id",
					Handler: useraccountsysadmin.QuerySysadminDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/useraccount/sysadmin/list",
					Handler: useraccountsysadmin.QuerySysadminListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/useraccount/sysadmin/page",
					Handler: useraccountsysadmin.QuerySysadminPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/useraccount/sysadmin",
					Handler: useraccountsysadmin.CreateSysadminHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/useraccount/sysadmin/:id",
					Handler: useraccountsysadmin.UpdateSysadminHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/useraccount/sysadmin/:id",
					Handler: useraccountsysadmin.DeleteSysadminHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/useraccount/user/:id",
				Handler: useraccountuser.QueryUserDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/useraccount/user/list",
				Handler: useraccountuser.QueryUserListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/useraccount/user/page",
				Handler: useraccountuser.QueryUserPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/useraccount/user",
				Handler: useraccountuser.CreateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/useraccount/user/:id",
				Handler: useraccountuser.UpdateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/useraccount/user/:id",
				Handler: useraccountuser.DeleteUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/system/api/v1"),
	)
}
