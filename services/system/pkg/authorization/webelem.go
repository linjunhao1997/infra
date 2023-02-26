package authorization

import "time"

type WebElem struct {
	Id          int64     `json:"id" gorm:"column:id;primary_key" redis:"id"`
	Name        string    `json:"name" gorm:"column:name" redis:"name"`
	Type        string    `json:"type" gorm:"column:type" redis:"type"`
	Code        string    `json:"code" gorm:"column:code" redis:"code"`
	IsOnline    bool      `json:"isOnline" gorm:"column:is_online" redis:"isOnline"`
	Description string    `json:"description" gorm:"column:description" redis:"description"`
	Ctime       time.Time `json:"createTime" gorm:"column:ctime;autoCreateTime:milli" redis:"mtime"`
	Mtime       time.Time `json:"updateTime" gorm:"column:mtime;autoUpdateTIme:milli" redis:"ctime"`
	CreatorUid  string    `json:"creatorUid" gorm:"column:creator_uid" redis:"creatorUid"`
	ModifierUid string    `json:"modifierUid" gorm:"column:modifier_uid" redis:"modifierUid"`
}

func (elem *WebElem) TableName() string {
	return "authorization_webelem"
}
