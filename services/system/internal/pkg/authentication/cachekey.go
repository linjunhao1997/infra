package authentication

import (
	"fmt"
)

type AuthVerKey string

type UserSessionKey string

func BuildUserSessionKey(authToken string) UserSessionKey {
	return (UserSessionKey)(fmt.Sprintf(`system:userSession:%s`, authToken))
}

func BuildAuthVerKey(gid string) AuthVerKey {
	return (AuthVerKey)(fmt.Sprintf(`system:auth:%s:ver`, gid))
}
