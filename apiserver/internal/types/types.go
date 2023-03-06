// Code generated by goctl. DO NOT EDIT.
package types

type CommonIdData struct {
	Id int64 `json:"id"`
}

type CommonPageData struct {
	CurrentPage int64       `json:"currentPage"`
	PageSize    int64       `json:"pageSize"`
	Total       int64       `json:"total"`
	List        interface{} `json:"list"`
}

type CommonResp struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Err  string      `json:"err"`
	Data interface{} `json:"data"`
}

type CommonPageResp struct {
	Code uint32         `json:"code"`
	Msg  string         `json:"msg"`
	Err  string         `json:"err"`
	Data CommonPageData `json:"data"`
}

type GetAuthTokenReq struct {
	UserType         string            `json:"userType"`
	InternalAuthMode *InternalAuthMode `json:"internalAuthMode"`
	SsoAuthMode      *SSOAuthMode      `json:"ssoAuthMode"`
	Oauth2AuthMode   *Oauth2AuthMode   `json:"oauth2AuthMode"`
}

type InternalAuthMode struct {
	Uid   *string `json:"uid,optional"`
	Email *string `json:"email,optional"`
	Pwd   string  `json:"pwd"`
}

type SSOAuthMode struct {
	IAM   *string `json:"iam"`
	IDAAS *string `json:"idaas"`
}

type Oauth2AuthMode struct {
	Code     string `json:"code"`
	ClientId string `json:"clientId"`
}

type GetAuthTokenResp struct {
	Code uint32           `json:"code"`
	Msg  string           `json:"msg"`
	Err  string           `json:"err"`
	Data GetAuthTokenData `json:"data"`
}

type GetAuthTokenData struct {
	AuthToken string `json:"authToken"`
}

type InternalAuthReq struct {
	UserType string  `json:"userType"`
	Uid      *string `json:"uid,optional"`
	Email    *string `json:"email,optional"`
	Pwd      string  `json:"pwd"`
}

type WebApiData struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Url         string `json:"url"`
	IsOnline    bool   `json:"isOnline"`
	Description string `json:"description"`
	Ctime       int64  `json:"createTime"`
	Mtime       int64  `json:"updateTime"`
	CreatorUid  string `json:"creatorUid"`
	ModifierUid string `json:"modifierUid"`
}

type QueryWebApiDetailReq struct {
	Id int64 `path:"id"`
}

type QueryWebApiDetailResp struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Err  string      `json:"err"`
	Data *WebApiData `json:"data"`
}

type QueryWebApiListReq struct {
	Code     *string `json:"code"`
	Url      *string `json:"url"`
	IsOnline *bool   `json:"isOnline"`
}

type QueryWebApiListResp struct {
	Code uint32        `json:"code"`
	Msg  string        `json:"msg"`
	Err  string        `json:"err"`
	Data []*WebApiData `json:"data"`
}

type QueryWebApiPageReq struct {
	CurrentPage int64   `json:"currentPage"`
	PageSize    int64   `json:"pageSize"`
	Code        *string `json:"code"`
	Url         *string `json:"url"`
	IsOnline    *bool   `json:"isOnline"`
}

type WebApiPage struct {
	CurrentPage int64         `json:"currentPage"`
	PageSize    int64         `json:"pageSize"`
	Total       int64         `json:"total"`
	List        []*WebApiData `json:"list"`
}

type QueryWebApiPageResp struct {
	Code uint32     `json:"code"`
	Msg  string     `json:"msg"`
	Err  string     `json:"err"`
	Data WebApiPage `json:"data"`
}

type CreateWebApiReq struct {
	Code        string `json:"code"`
	Url         string `json:"url"`
	IsOnline    bool   `json:"isOnline"`
	Description string `json:"description"`
}

type UpdateWebApiReq struct {
	Id          int64   `path:"id"`
	Code        *string `json:"code"`
	Url         *string `json:"url"`
	IsOnline    *bool   `json:"isOnline"`
	Description *string `json:"description"`
}

type UpsertWebApiReq struct {
	Id int64 `path:"id"`
}

type RemoveWebApiReq struct {
	Id int64 `path:"id"`
}

type DeleteWebApiReq struct {
	Id int64 `path:"id"`
}

type AdminData struct {
	Name     string      `json:"name"`
	Pwd      string      `json:"pwd"`
	Disabled bool        `json:"disabled"`
	Ctime    int64       `json:"ctime"`
	Mtime    int64       `json:"mtime"`
	Roles    []*RoleData `json:"roles"`
}

type QueryAdminDetailReq struct {
	Id int64 `path:"id"`
}

type QueryAdminDetailResp struct {
	Code uint32     `json:"code"`
	Msg  string     `json:"msg"`
	Err  string     `json:"err"`
	Data *AdminData `json:"data"`
}

type QueryAdminListReq struct {
	StartCtime *int64 `json:"startCtime"`
	EndCtime   *int64 `json:"endCtime"`
	CtimeDesc  bool   `json:"ctimeDesc"`
}

type QueryAdminListResp struct {
	Code uint32       `json:"code"`
	Msg  string       `json:"msg"`
	Err  string       `json:"err"`
	Data []*AdminData `json:"data"`
}

type QueryAdminPageReq struct {
	CurrentPage int64  `json:"currentPage"`
	PageSize    int64  `json:"pageSize"`
	StartCtime  *int64 `json:"startCtime"`
	EndCtime    *int64 `json:"endCtime"`
	CtimeDesc   bool   `json:"ctimeDesc"`
}

type AdminPage struct {
	CurrentPage int64        `json:"currentPage"`
	PageSize    int64        `json:"pageSize"`
	Total       int64        `json:"total"`
	List        []*AdminData `json:"list"`
}

type QueryAdminPageResp struct {
	Code uint32    `json:"code"`
	Msg  string    `json:"msg"`
	Err  string    `json:"err"`
	Data AdminPage `json:"data"`
}

type CreateAdminReq struct {
	Name     string `json:"name"`
	Pwd      string `json:"pwd"`
	Disabled bool   `json:"disabled"`
}

type UpdateAdminReq struct {
	Id       int64   `path:"id"`
	Name     *string `json:"name"`
	Pwd      *string `json:"pwd"`
	Disabled *bool   `json:"disabled"`
}

type UpsertAdminReq struct {
	Id int64 `path:"id"`
}

type RemoveAdminReq struct {
	Id int64 `path:"id"`
}

type DeleteAdminReq struct {
	Id int64 `path:"id"`
}

type RoleData struct {
	Name        string `json:"name"`
	ForUserType string `json:"forUserType"`
	Disabled    bool   `json:"dsiabled"`
	Ctime       int64  `json:"ctime"`
	Mtime       int64  `json:"mtime"`
}

type QueryRoleDetailReq struct {
	Id int64 `path:"id"`
}

type QueryRoleDetailResp struct {
	Code uint32    `json:"code"`
	Msg  string    `json:"msg"`
	Err  string    `json:"err"`
	Data *RoleData `json:"data"`
}

type QueryRoleListReq struct {
	Name       *string `json:"name"`
	StartCtime *int64  `json:"startCtime"`
	EndCtime   *int64  `json:"endCtime"`
}

type QueryRoleListResp struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Err  string      `json:"err"`
	Data []*RoleData `json:"data"`
}

type QueryRolePageReq struct {
	CurrentPage int64   `json:"currentPage"`
	PageSize    int64   `json:"pageSize"`
	Name        *string `json:"name"`
	StartCtime  *int64  `json:"startCtime"`
	EndCtime    *int64  `json:"endCtime"`
}

type RolePage struct {
	CurrentPage int64       `json:"currentPage"`
	PageSize    int64       `json:"pageSize"`
	Total       int64       `json:"total"`
	List        []*RoleData `json:"list"`
}

type QueryRolePageResp struct {
	Code uint32   `json:"code"`
	Msg  string   `json:"msg"`
	Err  string   `json:"err"`
	Data RolePage `json:"data"`
}

type CreateRoleReq struct {
	Name        string `json:"name"`
	ForUserType string `json:"forUserType"`
	Disabled    bool   `json:"dsiabled"`
}

type UpdateRoleReq struct {
	Id       int64   `path:"id"`
	Name     *string `json:"name"`
	Disabled *bool   `json:"dsiabled"`
}

type UpsertRoleReq struct {
	Id int64 `path:"id"`
}

type RemoveRoleReq struct {
	Id int64 `path:"id"`
}

type DeleteRoleReq struct {
	Id int64 `path:"id"`
}

type SysadminData struct {
	Name     string      `json:"name"`
	Pwd      string      `json:"pwd"`
	Disabled int64       `json:"disabled"`
	Ctime    int64       `json:"ctime"`
	Mtime    int64       `json:"mtime"`
	Roles    []*RoleData `json:"roles"`
}

type QuerySysadminDetailReq struct {
	Id int64 `path:"id"`
}

type QuerySysadminDetailResp struct {
	Code uint32        `json:"code"`
	Msg  string        `json:"msg"`
	Err  string        `json:"err"`
	Data *SysadminData `json:"data"`
}

type QuerySysadminListReq struct {
	Uid        *string  `json:"uid"`
	Name       *string  `json:"name"`
	Disabled   []*int64 `json:"disabled"`
	StartCtime *int64   `json:"startCtime"`
	EndCtime   *int64   `json:"endCtime"`
	CtimeDesc  bool     `json:"ctimeDesc"`
}

type QuerySysadminListResp struct {
	Code uint32          `json:"code"`
	Msg  string          `json:"msg"`
	Err  string          `json:"err"`
	Data []*SysadminData `json:"data"`
}

type QuerySysadminPageReq struct {
	CurrentPage int64    `json:"currentPage"`
	PageSize    int64    `json:"pageSize"`
	Uid         *string  `json:"uid"`
	Name        *string  `json:"name"`
	Disabled    []*int64 `json:"disabled"`
	StartCtime  *int64   `json:"startCtime"`
	EndCtime    *int64   `json:"endCtime"`
	CtimeDesc   bool     `json:"ctimeDesc"`
}

type SysadminPage struct {
	CurrentPage int64           `json:"currentPage"`
	PageSize    int64           `json:"pageSize"`
	Total       int64           `json:"total"`
	List        []*SysadminData `json:"list"`
}

type QuerySysadminPageResp struct {
	Code uint32       `json:"code"`
	Msg  string       `json:"msg"`
	Err  string       `json:"err"`
	Data SysadminPage `json:"data"`
}

type CreateSysadminReq struct {
	Name     string `json:"name"`
	Pwd      string `json:"pwd"`
	Disabled int64  `json:"disabled"`
}

type UpdateSysadminReq struct {
	Id       int64   `path:"id"`
	Name     *string `json:"name"`
	Pwd      *string `json:"pwd"`
	Disabled *int64  `json:"disabled"`
}

type UpsertSysadminReq struct {
	Id int64 `path:"id"`
}

type RemoveSysadminReq struct {
	Id int64 `path:"id"`
}

type DeleteSysadminReq struct {
	Id int64 `path:"id"`
}

type AuthenticateSysadmin struct {
	Uid   *string `json:"uid"`
	Email *string `json:"email"`
	Pwd   string  `json:"pwd"`
}

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Ctime int64  `json:"ctime"`
	Mtime int64  `json:"mtime"`
}

type QueryUserDetailReq struct {
	Id int64 `path:"id"`
}

type QueryUserDetailResp struct {
	Code uint32    `json:"code"`
	Msg  string    `json:"msg"`
	Err  string    `json:"err"`
	Data *UserData `json:"data"`
}

type QueryUserListReq struct {
	StartCtime *int64 `json:"startCtime"`
	EndCtime   *int64 `json:"endCtime"`
	CtimeDesc  bool   `json:"ctimeDesc"`
}

type QueryUserListResp struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Err  string      `json:"err"`
	Data []*UserData `json:"data"`
}

type QueryUserPageReq struct {
	CurrentPage int64  `json:"currentPage"`
	PageSize    int64  `json:"pageSize"`
	StartCtime  *int64 `json:"startCtime"`
	EndCtime    *int64 `json:"endCtime"`
	CtimeDesc   bool   `json:"ctimeDesc"`
}

type UserPage struct {
	CurrentPage int64       `json:"currentPage"`
	PageSize    int64       `json:"pageSize"`
	Total       int64       `json:"total"`
	List        []*UserData `json:"list"`
}

type QueryUserPageResp struct {
	Code uint32   `json:"code"`
	Msg  string   `json:"msg"`
	Err  string   `json:"err"`
	Data UserPage `json:"data"`
}

type CreateUserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UpdateUserReq struct {
	Id       int64   `path:"id"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Disabled *bool   `json:"disabled"`
}

type UpsertUserReq struct {
	Id int64 `path:"id"`
}

type RemoveUserReq struct {
	Id int64 `path:"id"`
}

type DeleteUserReq struct {
	Id int64 `path:"id"`
}
