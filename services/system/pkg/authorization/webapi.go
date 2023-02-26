package authorization

import "time"

type WebApi struct {
	Id          int64     `json:"id" gorm:"column:id;primary_key" redis:"id" gozero:"queryresp"`
	Code        string    `json:"code" gorm:"column:code" redis:"code" gozero:"createreq;updatereq;queryreq;queryresp"`
	Url         string    `json:"url" gorm:"column:url" redis:"url" gozero:"createreq;updatereq;queryreq;queryresp"`
	IsOnline    bool      `json:"isOnline" gorm:"column:is_online" redis:"isOnline" gozero:"createreq;updatereq;queryreq;queryresp"`
	Description string    `json:"description" gorm:"column:description" redis:"description" gozero:"createreq;updatereq;queryresp"`
	Ctime       time.Time `json:"createTime" gorm:"column:ctime;autoCreateTime:milli" redis:"mtime" gozero:"queryresp:int64"`
	Mtime       time.Time `json:"updateTime" gorm:"column:mtime;autoUpdateTIme:milli" redis:"ctime" gozero:"queryresp:int64"`
	CreatorUid  string    `json:"creatorUid" gorm:"column:creator_uid" redis:"creatorUid" gozero:"queryresp"`
	ModifierUid string    `json:"modifierUid" gorm:"column:modifier_uid" redis:"modifierUid" gozero:"queryresp"`
}

func (api *WebApi) TableName() string {
	return "authorization_webapi"
}
