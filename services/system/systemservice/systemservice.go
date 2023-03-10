// Code generated by goctl. DO NOT EDIT!
// Source: system.proto

package systemservice

import (
	"context"

	"infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AdminData               = v1.AdminData
	AdminPage               = v1.AdminPage
	CommonIdData            = v1.CommonIdData
	CommonIdDataResp        = v1.CommonIdDataResp
	CommonPageData          = v1.CommonPageData
	CommonPageResp          = v1.CommonPageResp
	CommonResp              = v1.CommonResp
	CreateAdminReq          = v1.CreateAdminReq
	CreateRoleReq           = v1.CreateRoleReq
	CreateSysadminReq       = v1.CreateSysadminReq
	CreateUserReq           = v1.CreateUserReq
	CreateWebApiReq         = v1.CreateWebApiReq
	CreateWebElemReq        = v1.CreateWebElemReq
	DeleteAdminReq          = v1.DeleteAdminReq
	DeleteRoleReq           = v1.DeleteRoleReq
	DeleteSysadminReq       = v1.DeleteSysadminReq
	DeleteUserReq           = v1.DeleteUserReq
	DeleteWebApiReq         = v1.DeleteWebApiReq
	DeleteWebElemReq        = v1.DeleteWebElemReq
	GetAuthTokenData        = v1.GetAuthTokenData
	GetAuthTokenReq         = v1.GetAuthTokenReq
	GetAuthTokenResp        = v1.GetAuthTokenResp
	InternalAuthMode        = v1.InternalAuthMode
	Oauth2AuthMode          = v1.Oauth2AuthMode
	QueryAdminDetailReq     = v1.QueryAdminDetailReq
	QueryAdminDetailResp    = v1.QueryAdminDetailResp
	QueryAdminListReq       = v1.QueryAdminListReq
	QueryAdminListResp      = v1.QueryAdminListResp
	QueryAdminPageReq       = v1.QueryAdminPageReq
	QueryAdminPageResp      = v1.QueryAdminPageResp
	QueryRoleDetailReq      = v1.QueryRoleDetailReq
	QueryRoleDetailResp     = v1.QueryRoleDetailResp
	QueryRoleListReq        = v1.QueryRoleListReq
	QueryRoleListResp       = v1.QueryRoleListResp
	QueryRolePageReq        = v1.QueryRolePageReq
	QueryRolePageResp       = v1.QueryRolePageResp
	QuerySysadminDetailReq  = v1.QuerySysadminDetailReq
	QuerySysadminDetailResp = v1.QuerySysadminDetailResp
	QuerySysadminListReq    = v1.QuerySysadminListReq
	QuerySysadminListResp   = v1.QuerySysadminListResp
	QuerySysadminPageReq    = v1.QuerySysadminPageReq
	QuerySysadminPageResp   = v1.QuerySysadminPageResp
	QueryUserDetailReq      = v1.QueryUserDetailReq
	QueryUserDetailResp     = v1.QueryUserDetailResp
	QueryUserListReq        = v1.QueryUserListReq
	QueryUserListResp       = v1.QueryUserListResp
	QueryUserPageReq        = v1.QueryUserPageReq
	QueryUserPageResp       = v1.QueryUserPageResp
	QueryWebApiDetailReq    = v1.QueryWebApiDetailReq
	QueryWebApiDetailResp   = v1.QueryWebApiDetailResp
	QueryWebApiListReq      = v1.QueryWebApiListReq
	QueryWebApiListResp     = v1.QueryWebApiListResp
	QueryWebApiPageReq      = v1.QueryWebApiPageReq
	QueryWebApiPageResp     = v1.QueryWebApiPageResp
	QueryWebElemDetailReq   = v1.QueryWebElemDetailReq
	QueryWebElemDetailResp  = v1.QueryWebElemDetailResp
	QueryWebElemListReq     = v1.QueryWebElemListReq
	QueryWebElemListResp    = v1.QueryWebElemListResp
	QueryWebElemPageReq     = v1.QueryWebElemPageReq
	QueryWebElemPageResp    = v1.QueryWebElemPageResp
	RemoveAdminReq          = v1.RemoveAdminReq
	RemoveRoleReq           = v1.RemoveRoleReq
	RemoveSysadminReq       = v1.RemoveSysadminReq
	RemoveUserReq           = v1.RemoveUserReq
	RemoveWebApiReq         = v1.RemoveWebApiReq
	RemoveWebElemReq        = v1.RemoveWebElemReq
	RoleData                = v1.RoleData
	RolePage                = v1.RolePage
	SSOAuthMode             = v1.SSOAuthMode
	SysadminData            = v1.SysadminData
	SysadminPage            = v1.SysadminPage
	UpdateAdminReq          = v1.UpdateAdminReq
	UpdateRoleReq           = v1.UpdateRoleReq
	UpdateSysadminReq       = v1.UpdateSysadminReq
	UpdateUserReq           = v1.UpdateUserReq
	UpdateWebApiReq         = v1.UpdateWebApiReq
	UpdateWebElemReq        = v1.UpdateWebElemReq
	UpsertAdminReq          = v1.UpsertAdminReq
	UpsertRoleReq           = v1.UpsertRoleReq
	UpsertSysadminReq       = v1.UpsertSysadminReq
	UpsertUserReq           = v1.UpsertUserReq
	UpsertWebApiReq         = v1.UpsertWebApiReq
	UpsertWebElemReq        = v1.UpsertWebElemReq
	UserData                = v1.UserData
	UserPage                = v1.UserPage
	WebApiData              = v1.WebApiData
	WebApiPage              = v1.WebApiPage
	WebElemData             = v1.WebElemData
	WebElemPage             = v1.WebElemPage

	SystemService interface {
		QueryRoleDetail(ctx context.Context, in *QueryRoleDetailReq, opts ...grpc.CallOption) (*QueryRoleDetailResp, error)
		QueryRoleList(ctx context.Context, in *QueryRoleListReq, opts ...grpc.CallOption) (*QueryRoleListResp, error)
		QueryRolePage(ctx context.Context, in *QueryRolePageReq, opts ...grpc.CallOption) (*QueryRolePageResp, error)
		CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CommonIdDataResp, error)
		UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*CommonResp, error)
		DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*CommonResp, error)
		QueryUserDetail(ctx context.Context, in *QueryUserDetailReq, opts ...grpc.CallOption) (*QueryUserDetailResp, error)
		QueryUserList(ctx context.Context, in *QueryUserListReq, opts ...grpc.CallOption) (*QueryUserListResp, error)
		QueryUserPage(ctx context.Context, in *QueryUserPageReq, opts ...grpc.CallOption) (*QueryUserPageResp, error)
		CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CommonIdDataResp, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*CommonResp, error)
		DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*CommonResp, error)
		QueryAdminDetail(ctx context.Context, in *QueryAdminDetailReq, opts ...grpc.CallOption) (*QueryAdminDetailResp, error)
		QueryAdminList(ctx context.Context, in *QueryAdminListReq, opts ...grpc.CallOption) (*QueryAdminListResp, error)
		QueryAdminPage(ctx context.Context, in *QueryAdminPageReq, opts ...grpc.CallOption) (*QueryAdminPageResp, error)
		CreateAdmin(ctx context.Context, in *CreateAdminReq, opts ...grpc.CallOption) (*CommonIdDataResp, error)
		UpdateAdmin(ctx context.Context, in *UpdateAdminReq, opts ...grpc.CallOption) (*CommonResp, error)
		DeleteAdmin(ctx context.Context, in *DeleteAdminReq, opts ...grpc.CallOption) (*CommonResp, error)
		QuerySysadminDetail(ctx context.Context, in *QuerySysadminDetailReq, opts ...grpc.CallOption) (*QuerySysadminDetailResp, error)
		QuerySysadminList(ctx context.Context, in *QuerySysadminListReq, opts ...grpc.CallOption) (*QuerySysadminListResp, error)
		QuerySysadminPage(ctx context.Context, in *QuerySysadminPageReq, opts ...grpc.CallOption) (*QuerySysadminPageResp, error)
		CreateSysadmin(ctx context.Context, in *CreateSysadminReq, opts ...grpc.CallOption) (*CommonIdDataResp, error)
		UpdateSysadmin(ctx context.Context, in *UpdateSysadminReq, opts ...grpc.CallOption) (*CommonResp, error)
		DeleteSysadmin(ctx context.Context, in *DeleteSysadminReq, opts ...grpc.CallOption) (*CommonResp, error)
		GetAuthToken(ctx context.Context, in *GetAuthTokenReq, opts ...grpc.CallOption) (*GetAuthTokenResp, error)
		QueryWebElemDetail(ctx context.Context, in *QueryWebElemDetailReq, opts ...grpc.CallOption) (*QueryWebElemDetailResp, error)
		QueryWebElemList(ctx context.Context, in *QueryWebElemListReq, opts ...grpc.CallOption) (*QueryWebElemListResp, error)
		QueryWebElemPage(ctx context.Context, in *QueryWebElemPageReq, opts ...grpc.CallOption) (*QueryWebElemPageResp, error)
		CreateWebElem(ctx context.Context, in *CreateWebElemReq, opts ...grpc.CallOption) (*CommonIdDataResp, error)
		UpdateWebElem(ctx context.Context, in *UpdateWebElemReq, opts ...grpc.CallOption) (*CommonResp, error)
		DeleteWebElem(ctx context.Context, in *DeleteWebElemReq, opts ...grpc.CallOption) (*CommonResp, error)
		QueryWebApiDetail(ctx context.Context, in *QueryWebApiDetailReq, opts ...grpc.CallOption) (*QueryWebApiDetailResp, error)
		QueryWebApiList(ctx context.Context, in *QueryWebApiListReq, opts ...grpc.CallOption) (*QueryWebApiListResp, error)
		QueryWebApiPage(ctx context.Context, in *QueryWebApiPageReq, opts ...grpc.CallOption) (*QueryWebApiPageResp, error)
		CreateWebApi(ctx context.Context, in *CreateWebApiReq, opts ...grpc.CallOption) (*CommonIdDataResp, error)
		UpdateWebApi(ctx context.Context, in *UpdateWebApiReq, opts ...grpc.CallOption) (*CommonResp, error)
		DeleteWebApi(ctx context.Context, in *DeleteWebApiReq, opts ...grpc.CallOption) (*CommonResp, error)
	}

	defaultSystemService struct {
		cli zrpc.Client
	}
)

func NewSystemService(cli zrpc.Client) SystemService {
	return &defaultSystemService{
		cli: cli,
	}
}

func (m *defaultSystemService) QueryRoleDetail(ctx context.Context, in *QueryRoleDetailReq, opts ...grpc.CallOption) (*QueryRoleDetailResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryRoleDetail(ctx, in, opts...)
}

func (m *defaultSystemService) QueryRoleList(ctx context.Context, in *QueryRoleListReq, opts ...grpc.CallOption) (*QueryRoleListResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryRoleList(ctx, in, opts...)
}

func (m *defaultSystemService) QueryRolePage(ctx context.Context, in *QueryRolePageReq, opts ...grpc.CallOption) (*QueryRolePageResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryRolePage(ctx, in, opts...)
}

func (m *defaultSystemService) CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CommonIdDataResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.CreateRole(ctx, in, opts...)
}

func (m *defaultSystemService) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

func (m *defaultSystemService) DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

func (m *defaultSystemService) QueryUserDetail(ctx context.Context, in *QueryUserDetailReq, opts ...grpc.CallOption) (*QueryUserDetailResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryUserDetail(ctx, in, opts...)
}

func (m *defaultSystemService) QueryUserList(ctx context.Context, in *QueryUserListReq, opts ...grpc.CallOption) (*QueryUserListResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryUserList(ctx, in, opts...)
}

func (m *defaultSystemService) QueryUserPage(ctx context.Context, in *QueryUserPageReq, opts ...grpc.CallOption) (*QueryUserPageResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryUserPage(ctx, in, opts...)
}

func (m *defaultSystemService) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CommonIdDataResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultSystemService) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultSystemService) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

func (m *defaultSystemService) QueryAdminDetail(ctx context.Context, in *QueryAdminDetailReq, opts ...grpc.CallOption) (*QueryAdminDetailResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryAdminDetail(ctx, in, opts...)
}

func (m *defaultSystemService) QueryAdminList(ctx context.Context, in *QueryAdminListReq, opts ...grpc.CallOption) (*QueryAdminListResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryAdminList(ctx, in, opts...)
}

func (m *defaultSystemService) QueryAdminPage(ctx context.Context, in *QueryAdminPageReq, opts ...grpc.CallOption) (*QueryAdminPageResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryAdminPage(ctx, in, opts...)
}

func (m *defaultSystemService) CreateAdmin(ctx context.Context, in *CreateAdminReq, opts ...grpc.CallOption) (*CommonIdDataResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.CreateAdmin(ctx, in, opts...)
}

func (m *defaultSystemService) UpdateAdmin(ctx context.Context, in *UpdateAdminReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.UpdateAdmin(ctx, in, opts...)
}

func (m *defaultSystemService) DeleteAdmin(ctx context.Context, in *DeleteAdminReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.DeleteAdmin(ctx, in, opts...)
}

func (m *defaultSystemService) QuerySysadminDetail(ctx context.Context, in *QuerySysadminDetailReq, opts ...grpc.CallOption) (*QuerySysadminDetailResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QuerySysadminDetail(ctx, in, opts...)
}

func (m *defaultSystemService) QuerySysadminList(ctx context.Context, in *QuerySysadminListReq, opts ...grpc.CallOption) (*QuerySysadminListResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QuerySysadminList(ctx, in, opts...)
}

func (m *defaultSystemService) QuerySysadminPage(ctx context.Context, in *QuerySysadminPageReq, opts ...grpc.CallOption) (*QuerySysadminPageResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QuerySysadminPage(ctx, in, opts...)
}

func (m *defaultSystemService) CreateSysadmin(ctx context.Context, in *CreateSysadminReq, opts ...grpc.CallOption) (*CommonIdDataResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.CreateSysadmin(ctx, in, opts...)
}

func (m *defaultSystemService) UpdateSysadmin(ctx context.Context, in *UpdateSysadminReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.UpdateSysadmin(ctx, in, opts...)
}

func (m *defaultSystemService) DeleteSysadmin(ctx context.Context, in *DeleteSysadminReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.DeleteSysadmin(ctx, in, opts...)
}

func (m *defaultSystemService) GetAuthToken(ctx context.Context, in *GetAuthTokenReq, opts ...grpc.CallOption) (*GetAuthTokenResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.GetAuthToken(ctx, in, opts...)
}

func (m *defaultSystemService) QueryWebElemDetail(ctx context.Context, in *QueryWebElemDetailReq, opts ...grpc.CallOption) (*QueryWebElemDetailResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryWebElemDetail(ctx, in, opts...)
}

func (m *defaultSystemService) QueryWebElemList(ctx context.Context, in *QueryWebElemListReq, opts ...grpc.CallOption) (*QueryWebElemListResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryWebElemList(ctx, in, opts...)
}

func (m *defaultSystemService) QueryWebElemPage(ctx context.Context, in *QueryWebElemPageReq, opts ...grpc.CallOption) (*QueryWebElemPageResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryWebElemPage(ctx, in, opts...)
}

func (m *defaultSystemService) CreateWebElem(ctx context.Context, in *CreateWebElemReq, opts ...grpc.CallOption) (*CommonIdDataResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.CreateWebElem(ctx, in, opts...)
}

func (m *defaultSystemService) UpdateWebElem(ctx context.Context, in *UpdateWebElemReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.UpdateWebElem(ctx, in, opts...)
}

func (m *defaultSystemService) DeleteWebElem(ctx context.Context, in *DeleteWebElemReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.DeleteWebElem(ctx, in, opts...)
}

func (m *defaultSystemService) QueryWebApiDetail(ctx context.Context, in *QueryWebApiDetailReq, opts ...grpc.CallOption) (*QueryWebApiDetailResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryWebApiDetail(ctx, in, opts...)
}

func (m *defaultSystemService) QueryWebApiList(ctx context.Context, in *QueryWebApiListReq, opts ...grpc.CallOption) (*QueryWebApiListResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryWebApiList(ctx, in, opts...)
}

func (m *defaultSystemService) QueryWebApiPage(ctx context.Context, in *QueryWebApiPageReq, opts ...grpc.CallOption) (*QueryWebApiPageResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.QueryWebApiPage(ctx, in, opts...)
}

func (m *defaultSystemService) CreateWebApi(ctx context.Context, in *CreateWebApiReq, opts ...grpc.CallOption) (*CommonIdDataResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.CreateWebApi(ctx, in, opts...)
}

func (m *defaultSystemService) UpdateWebApi(ctx context.Context, in *UpdateWebApiReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.UpdateWebApi(ctx, in, opts...)
}

func (m *defaultSystemService) DeleteWebApi(ctx context.Context, in *DeleteWebApiReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := v1.NewSystemServiceClient(m.cli.Conn())
	return client.DeleteWebApi(ctx, in, opts...)
}
