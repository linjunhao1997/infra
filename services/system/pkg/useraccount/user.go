package useraccount

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type User struct {
	Id int64 `gorm:"column:id;primaryKey"`
	// gid为root账号的uid
	Gid       int64     `gorm:"column:gid" json:"gid"`
	Uid       string    `gorm:"column:uid" json:"uid" gozero:"queryresp"`
	Name      string    `gorm:"column:name" json:"name" gozero:"queryresp;createreq"`
	Email     string    `gorm:"column:email" json:"email" gozero:"queryresp;updatereq;createreq"`
	Phone     string    `gorm:"column:phone" json:"phone" gozero:"queryresp;updatereq;createreq"`
	Pwd       string    `gorm:"column:pwd"`
	Disabled  bool      `gorm:"column:disabled" json:"disabled" gozero:"queryresp;updatereq"`
	LastPatch string    `gorm:"column:last_patch"`
	Ver       int64     `gorm:"column:ver"`
	Ctime     time.Time `gorm:"column:ctime;autoCreateTime" json:"ctime" gozero:"queryreq;queryreqtype:*int64;queryreqscope;queryreqsort;queryresptype:int64"`
	Mtime     time.Time `gorm:"column:mtime;autoUpdateTime" json:"mtime" gozero:"queryresptype:int64"`
}

func (user *User) TableName() string {
	return "useraccount_user"
}

func (user *User) GenerateUid() *User {
	if user == nil {
		return user
	}

	if len(user.Uid) > 0 {
		return user
	}

	uid, _ := gonanoid.New()
	user.Uid = uid
	return user
}
