package authentication

import (
	"encoding/json"
	"errors"
)

var (
	errUserSessionIsNil = errors.New(`usersession is nil`)
)

type UserSession struct {
	Uid      string   `redis:"uid"`
	Gid      string   `redis:"gid"`
	UserType string   `redis:"userType"`
	Roles    []string `redis:"roles"`
	AuthVer  string   `redis:"authVer"`
}

func (u *UserSession) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *UserSession) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
