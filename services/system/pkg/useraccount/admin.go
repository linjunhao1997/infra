package useraccount

import (
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
)

type Admin struct {
	Id int64 `gorm:"column:id;primaryKey" json:"id" gozero:"queryresp"`
	// 数据隔离
	Gid string `gorm:"column:gid" json:"gid" gozero:"queryresp"`
	//禁止修改, 0开头的用户是root级别的后台用户
	Uid      string    `gorm:"column:uid" json:"uid" gozero:"queryresp"`
	Name     string    `gorm:"column:name" json:"name" gozero:"createreq;updatereq"`
	Pwd      string    `gorm:"column:pwd" json:"pwd" gozero:"createreq;updatereq"`
	Disabled bool      `gorm:"column:disabled" json:"disabled" gozero:"createreq;updatereq;queryresp"`
	Ver      int64     `gorm:"column:ver"`
	Ctime    time.Time `gorm:"column:ctime;autoCreateTime" json:"ctime" gozero:"queryreq;queryreqtype:*int64;queryreqscope;queryreqsort;queryresp;queryresptype:int64"`
	Mtime    time.Time `gorm:"column:mtime;autoUpdateTime" json:"mtime" gozero:"queryresp;queryresptype:int64"`
	Roles    Roles     `gorm:"many2many:useraccount_admin_m2m_role" json:"roles" gozero:"queryresp;queryresptype:[]*RoleData"`
}

func (admin *Admin) TableName() string {
	return "useraccount_admin"
}

func (admin *Admin) GenerateRootUid() *Admin {
	if admin == nil {
		return admin
	}

	if len(admin.Uid) > 0 {
		return admin
	}

	uid, _ := nanoid.New()
	admin.Uid = RootTag + uid
	return admin
}

func (admin *Admin) GenerateUid() *Admin {
	if admin == nil {
		return admin
	}

	if len(admin.Uid) > 0 {
		return admin
	}

	uid, _ := nanoid.New()
	admin.Uid = uid
	return admin
}

func (admin *Admin) IsRoot() bool {
	if len(admin.Uid) < RootUidLength {
		return false
	}

	if len(admin.Uid) > RootUidLength {
		return false
	}

	return admin.Uid[:1] == RootTag
}
