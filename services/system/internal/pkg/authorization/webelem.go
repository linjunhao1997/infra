package authorization

import "time"

type WebElem struct {
	Id          int64     `json:"id" gorm:"column:id;primary_key" redis:"id" gozero:"queryresp"`
	Name        string    `json:"name" gorm:"column:name" redis:"name" gozero:"createreq;updatereq;queryresp"`
	Type        string    `json:"type" gorm:"column:type" redis:"type" gozero:"createreq;queryresp"`
	Code        string    `json:"code" gorm:"column:code" redis:"code" gozero:"createreq;queryresp"`
	Disabled    bool      `json:"disabled" gorm:"column:disabled" redis:"disabled" gozero:"createreq;updatereq;queryresp"`
	Description string    `json:"description" gorm:"column:description" redis:"description" gozero:"createreq;updatereq;queryresp"`
	Ctime       time.Time `json:"createTime" gorm:"column:ctime;autoCreateTime:milli" redis:"mtime" gozero:"queryresp"`
	Mtime       time.Time `json:"updateTime" gorm:"column:mtime;autoUpdateTIme:milli" redis:"ctime" gozero:"queryresp"`
	CreatorUid  string    `json:"creatorUid" gorm:"column:creator_uid" redis:"creatorUid" gozero:"queryresp"`
	ModifierUid string    `json:"modifierUid" gorm:"column:modifier_uid" redis:"modifierUid" gozero:"queryresp"`
}

func (elem *WebElem) TableName() string {
	return "authorization_webelem"
}
