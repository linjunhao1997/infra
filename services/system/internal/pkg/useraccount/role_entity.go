package useraccount

import "time"

type Role struct {
	Id int64 `gorm:"column:id;primaryKey" json:"id,string" gozeroapi:"queryresp"`
	// gid为root账号的uid
	Gid         string    `gorm:"column:gid" json:"gid" gozero:"queryresp"`
	Name        string    `gorm:"column:name" json:"name" gozero:"queryreq;queryresp;createreq;updatereq"`
	Description string    `gorm:"column:description" json:"description" gozero:"queryresp;updatereq,createreq"`
	ForUserType string    `gorm:"column:for_user_type" json:"forUserType" gozero:"queryresp;createreq"`
	Disabled    bool      `gorm:"column:disabled" json:"dsiabled" gozero:"queryresp;createreq;updatereq"`
	Ver         int64     `gorm:"column:ver"`
	Ctime       time.Time `gorm:"column:ctime;autoCreateTime" json:"ctime" gozero:"queryreqtype:*int64;queryreqscope;queryresp;queryresptype:int64"`
	Mtime       time.Time `gorm:"column:mtime;autoUpdateTime" json:"mtime" gozero:"queryresp;queryresptype:int64"`
}

func (role *Role) TableName() string {
	return "useraccount_role"
}

type Roles []*Role

func (roles Roles) Names() []string {
	res := make([]string, 0, len(roles))
	for _, e := range roles {
		res = append(res, e.Name)
	}
	return res
}
