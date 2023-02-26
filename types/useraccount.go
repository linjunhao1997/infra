package types

const (
	UserAccountTypeSysUser = "sysuser"
	UserAccountTypeAdmin   = "admin"
	UserAccountTypeUser    = "user"
)

var UserAccountType = &userAccountType{
	SysUser: UserAccountTypeSysUser,
	Admin:   UserAccountTypeAdmin,
	User:    UserAccountTypeUser,
}

// UserAccount 用户类型
type userAccountType struct {
	// 系统 管理员
	SysUser string
	// 后台 管理员
	Admin string
	// 用户
	User string
}
