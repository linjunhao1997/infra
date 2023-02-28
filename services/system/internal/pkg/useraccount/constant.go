package useraccount

const (
	RootUidLength = 22
	RootTag       = "0"
)

const (
	UserTypeSysadmin = "sysadmin"
	UserTypeAdmin    = "admin"
)

func IsValidUserType(forUserType string) bool {
	ls := []string{UserTypeAdmin, UserTypeAdmin}
	for _, e := range ls {
		if forUserType == e {
			return true
		}
	}
	return false
}
