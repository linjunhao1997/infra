package errorx

var (
	// 微服务编号2位+微服务模块2位+4位错误码

	// 认证失败
	ErrAuthenticationFalied              = NewExpectError(10010001, "authentication failed")
	ErrAuthenticationInfoExpired         = NewExpectError(10010002, "authentication info expired")
	ErrAuthenticationAuthModeUnsupported = NewExpectError(10010003, "authentication authmode unsupported")
	ErrAuthenticationUserTypeUnsupported = NewExpectError(10010004, "authentication usertype unsupported")

	// 授权失败
	ErrAuthorizationFalied = NewExpectError(10020001, "authorization failed")
	// 无权限
	ErrPermissionDenied = NewExpectError(10020003, "permission denied")

	ErrSysadminNotFound     = NewExpectError(10030000, "sysadmin not found")
	ErrSysadminCreateFailed = NewExpectError(10030001, "sysadmin create failed")
	ErrSysadminUpdateFailed = NewExpectError(10030002, "sysadmin update failed")
	ErrSysadminDeleteFailed = NewExpectError(10030003, "sysadmin delete failed")
	ErrSysadminQueryFailed  = NewExpectError(10030004, "sysadmin query failed")
	ErrSysadminAuthFalied   = NewExpectError(10030005, "sysadmin authenticate failed")

	ErrAdminNotFound     = NewExpectError(10040000, "admin not found")
	ErrAdminCreateFailed = NewExpectError(10040001, "admin create failed")
	ErrAdminUpdateFailed = NewExpectError(10040002, "admin update failed")
	ErrAdminDeleteFailed = NewExpectError(10040003, "admin delete failed")
	ErrAdminQueryFailed  = NewExpectError(10040004, "admin query failed")

	ErrUserNotFound     = NewExpectError(10050000, "user not found")
	ErrUserCreateFailed = NewExpectError(10050001, "user create failed")
	ErrUserUpdateFailed = NewExpectError(10050002, "user update failed")
	ErrUserDeleteFailed = NewExpectError(10050003, "user delete failed")
	ErrUserQueryFailed  = NewExpectError(10050004, "user query failed")

	ErrRoleNotFound     = NewExpectError(10060000, "role not found")
	ErrRoleCreateFailed = NewExpectError(10060001, "role create failed")
	ErrRoleUpdateFailed = NewExpectError(10060002, "role update failed")
	ErrRoleDeleteFailed = NewExpectError(10060003, "role delete failed")
	ErrRoleQueryFailed  = NewExpectError(10060004, "role query failed")

	ErrWebApiNotFound     = NewExpectError(10070000, "webapi not found")
	ErrWebApiCreateFailed = NewExpectError(10070001, "webapi create failed")
	ErrWebApiUpdateFailed = NewExpectError(10070002, "webapi update failed")
	ErrWebApiDeleteFailed = NewExpectError(10070003, "webapi delete failed")
	ErrWebApiQueryFailed  = NewExpectError(10070004, "webapi query failed")
)
