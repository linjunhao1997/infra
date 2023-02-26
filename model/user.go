package model

import (
	"time"
)

// 无法删除，只能禁用
type User struct {
	Id int64
	//禁止修改
	Uid      string     `gorm:"column:uid"`
	Name     string     `gorm:"column:name"`
	Email    string     `gorm:"column:email"`
	Pwd      string     `gorm:"column:pwd"`
	Disabled bool       `gorm:"column:disabled"`
	Ver      int64      `gorm:"column:ver"`
	Ctime    time.Time  `gorm:"column:ctime;autoCreateTime"`
	Mtime    time.Time  `gorm:"column:mtime;autoUpdateTime"`
	Groups   UserGroups `gorm:"many2many:m2m_user_usergroup"`
}

func (user *User) TableName() string {
	return "user"
}

type Users []*User

// 无法删除，只能禁用
type UserGroup struct {
	Id int64
	// 禁止修改
	Gid         string    `gorm:"column:gid"`
	Name        string    `gorm:"column:name"`
	Type        string    `gorm:"column:type"`
	Disabled    bool      `gorm:"column:disabled"`
	Ver         int64     `gorm:"column:ver"`
	Ctime       time.Time `gorm:"column:ctime;autoCreateTime"`
	Mtime       time.Time `gorm:"column:mtime;autoUpdateTime"`
	OwnerUids   string    `gorm:"column:owner_uids"`
	CreatorUid  string    `gorm:"column:creator_uid"`
	ModifierUid string    `gorm:"column:modifier_uid"`
	Users       Users     `gorm:"many2many:m2m_user_usergroup"`
}

func (userGroup *UserGroup) TableName() string {
	return "usergroup"
}

type UserGroups []*UserGroup
