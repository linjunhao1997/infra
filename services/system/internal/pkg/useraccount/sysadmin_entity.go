package useraccount

import (
	"infra/services/system/internal/store"
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
)

// uid为0的系统用户拥有最高权限
type Sysadmin struct {
	Id int64 `gorm:"column:id;primaryKey" json:"id,string" gozero:"queryresp"`
	//禁止修改
	Gid      string    `gorm:"column:gid" json:"gid" gozero:"queryresp"`
	Uid      string    `gorm:"column:uid" json:"uid" gozero:"queryresp;queryreq"`
	Name     string    `gorm:"column:name" json:"name" gozero:"queryresp;createreq;updatereq;queryreq"`
	Pwd      string    `gorm:"column:pwd" json:"pwd"  gozero:"createreq;updatereq"`
	Disabled int64     `gorm:"column:disabled" json:"disabled" gozero:"queryresp;createreq;updatereq;queryreqmulti"`
	Ver      int64     `gorm:"column:ver"`
	Ctime    time.Time `gorm:"column:ctime;autoCreateTime" json:"ctime" gozero:"queryreqtype:*int64;queryreqscope;queryreqsort;queryresptype:int64"`
	Mtime    time.Time `gorm:"column:mtime;autoUpdateTime" json:"mtime" gozero:"queryresptype:int64"`
	Roles    Roles     `gorm:"many2many:useraccount_sysadmin_m2m_role" json:"roles" gozero:"queryresptype:[]*RoleData"`
}

func (user *Sysadmin) TableName() string {
	return "useraccount_sysadmin"
}

func (user *Sysadmin) IsRoot() bool {
	return user.Uid == RootTag
}

func (user *Sysadmin) GenerateUid() *Sysadmin {
	if user == nil {
		return user
	}

	if len(user.Uid) > 0 {
		return user
	}

	uid, _ := nanoid.New()
	user.Uid = RootTag + uid
	return user
}

func (user *Sysadmin) GetRoles() (Roles, error) {
	if err := store.GormDB.Model(Sysadmin{}).Preload("Roles", `id = ?`, user.Id).Find(user).Error; err != nil {
		return nil, err
	}
	return user.Roles, nil
}
