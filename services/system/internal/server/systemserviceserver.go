// Code generated by goctl. DO NOT EDIT!
// Source: system.proto

package server

import (
	"context"

	"infra/services/system/internal/logic"
	"infra/services/system/internal/svc"
	"infra/services/system/pb/v1"
)

type SystemServiceServer struct {
	svcCtx *svc.ServiceContext
	v1.UnimplementedSystemServiceServer
}

func NewSystemServiceServer(svcCtx *svc.ServiceContext) *SystemServiceServer {
	return &SystemServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *SystemServiceServer) QueryRoleDetail(ctx context.Context, in *v1.QueryRoleDetailReq) (*v1.QueryRoleDetailResp, error) {
	l := logic.NewQueryRoleDetailLogic(ctx, s.svcCtx)
	return l.QueryRoleDetail(in)
}

func (s *SystemServiceServer) QueryRoleList(ctx context.Context, in *v1.QueryRoleListReq) (*v1.QueryRoleListResp, error) {
	l := logic.NewQueryRoleListLogic(ctx, s.svcCtx)
	return l.QueryRoleList(in)
}

func (s *SystemServiceServer) QueryRolePage(ctx context.Context, in *v1.QueryRolePageReq) (*v1.QueryRolePageResp, error) {
	l := logic.NewQueryRolePageLogic(ctx, s.svcCtx)
	return l.QueryRolePage(in)
}

func (s *SystemServiceServer) CreateRole(ctx context.Context, in *v1.CreateRoleReq) (*v1.CommonIdDataResp, error) {
	l := logic.NewCreateRoleLogic(ctx, s.svcCtx)
	return l.CreateRole(in)
}

func (s *SystemServiceServer) UpdateRole(ctx context.Context, in *v1.UpdateRoleReq) (*v1.CommonResp, error) {
	l := logic.NewUpdateRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(in)
}

func (s *SystemServiceServer) DeleteRole(ctx context.Context, in *v1.DeleteRoleReq) (*v1.CommonResp, error) {
	l := logic.NewDeleteRoleLogic(ctx, s.svcCtx)
	return l.DeleteRole(in)
}

func (s *SystemServiceServer) QueryUserDetail(ctx context.Context, in *v1.QueryUserDetailReq) (*v1.QueryUserDetailResp, error) {
	l := logic.NewQueryUserDetailLogic(ctx, s.svcCtx)
	return l.QueryUserDetail(in)
}

func (s *SystemServiceServer) QueryUserList(ctx context.Context, in *v1.QueryUserListReq) (*v1.QueryUserListResp, error) {
	l := logic.NewQueryUserListLogic(ctx, s.svcCtx)
	return l.QueryUserList(in)
}

func (s *SystemServiceServer) QueryUserPage(ctx context.Context, in *v1.QueryUserPageReq) (*v1.QueryUserPageResp, error) {
	l := logic.NewQueryUserPageLogic(ctx, s.svcCtx)
	return l.QueryUserPage(in)
}

func (s *SystemServiceServer) CreateUser(ctx context.Context, in *v1.CreateUserReq) (*v1.CommonIdDataResp, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *SystemServiceServer) UpdateUser(ctx context.Context, in *v1.UpdateUserReq) (*v1.CommonResp, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *SystemServiceServer) DeleteUser(ctx context.Context, in *v1.DeleteUserReq) (*v1.CommonResp, error) {
	l := logic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *SystemServiceServer) QueryAdminDetail(ctx context.Context, in *v1.QueryAdminDetailReq) (*v1.QueryAdminDetailResp, error) {
	l := logic.NewQueryAdminDetailLogic(ctx, s.svcCtx)
	return l.QueryAdminDetail(in)
}

func (s *SystemServiceServer) QueryAdminList(ctx context.Context, in *v1.QueryAdminListReq) (*v1.QueryAdminListResp, error) {
	l := logic.NewQueryAdminListLogic(ctx, s.svcCtx)
	return l.QueryAdminList(in)
}

func (s *SystemServiceServer) QueryAdminPage(ctx context.Context, in *v1.QueryAdminPageReq) (*v1.QueryAdminPageResp, error) {
	l := logic.NewQueryAdminPageLogic(ctx, s.svcCtx)
	return l.QueryAdminPage(in)
}

func (s *SystemServiceServer) CreateAdmin(ctx context.Context, in *v1.CreateAdminReq) (*v1.CommonIdDataResp, error) {
	l := logic.NewCreateAdminLogic(ctx, s.svcCtx)
	return l.CreateAdmin(in)
}

func (s *SystemServiceServer) UpdateAdmin(ctx context.Context, in *v1.UpdateAdminReq) (*v1.CommonResp, error) {
	l := logic.NewUpdateAdminLogic(ctx, s.svcCtx)
	return l.UpdateAdmin(in)
}

func (s *SystemServiceServer) DeleteAdmin(ctx context.Context, in *v1.DeleteAdminReq) (*v1.CommonResp, error) {
	l := logic.NewDeleteAdminLogic(ctx, s.svcCtx)
	return l.DeleteAdmin(in)
}

func (s *SystemServiceServer) QuerySysadminDetail(ctx context.Context, in *v1.QuerySysadminDetailReq) (*v1.QuerySysadminDetailResp, error) {
	l := logic.NewQuerySysadminDetailLogic(ctx, s.svcCtx)
	return l.QuerySysadminDetail(in)
}

func (s *SystemServiceServer) QuerySysadminList(ctx context.Context, in *v1.QuerySysadminListReq) (*v1.QuerySysadminListResp, error) {
	l := logic.NewQuerySysadminListLogic(ctx, s.svcCtx)
	return l.QuerySysadminList(in)
}

func (s *SystemServiceServer) QuerySysadminPage(ctx context.Context, in *v1.QuerySysadminPageReq) (*v1.QuerySysadminPageResp, error) {
	l := logic.NewQuerySysadminPageLogic(ctx, s.svcCtx)
	return l.QuerySysadminPage(in)
}

func (s *SystemServiceServer) CreateSysadmin(ctx context.Context, in *v1.CreateSysadminReq) (*v1.CommonIdDataResp, error) {
	l := logic.NewCreateSysadminLogic(ctx, s.svcCtx)
	return l.CreateSysadmin(in)
}

func (s *SystemServiceServer) UpdateSysadmin(ctx context.Context, in *v1.UpdateSysadminReq) (*v1.CommonResp, error) {
	l := logic.NewUpdateSysadminLogic(ctx, s.svcCtx)
	return l.UpdateSysadmin(in)
}

func (s *SystemServiceServer) DeleteSysadmin(ctx context.Context, in *v1.DeleteSysadminReq) (*v1.CommonResp, error) {
	l := logic.NewDeleteSysadminLogic(ctx, s.svcCtx)
	return l.DeleteSysadmin(in)
}

func (s *SystemServiceServer) GetAuthToken(ctx context.Context, in *v1.GetAuthTokenReq) (*v1.GetAuthTokenResp, error) {
	l := logic.NewGetAuthTokenLogic(ctx, s.svcCtx)
	return l.GetAuthToken(in)
}
