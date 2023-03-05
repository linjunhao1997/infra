package types

const (
	UserAccountTypeSysadmin = "sysadmin"
	UserAccountTypeAdmin    = "admin"
	UserAccountTypeUser     = "user"
)

var UserAccountType = &userAccountType{
	Sysadmin: UserAccountTypeSysadmin,
	Admin:    UserAccountTypeAdmin,
	User:     UserAccountTypeUser,
}

// UserAccount 用户类型
type userAccountType struct {
	// 系统 管理员
	Sysadmin string
	// 后台 管理员
	Admin string
	// 用户
	User string
}
