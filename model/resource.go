package model

import (
	"time"

	"gorm.io/gorm"
)

// 资源型通用模型
type Resource struct {
	// 主键id
	Id int64 `gorm:"column:id;primaryKey"`
	// 资源名
	Name string `gorm:"column:name"`
	// 固定static、安全security的id（对此id涉及资源进行固定场景的需求开发）
	Sid string `gorm:"column:sid"`
	// 创建人Uid
	CreatorUid string `gorm:"column:creator_uid"`
	// 修改人uid
	ModifierUid string `gorm:"column:modifier_uid"`
	// 读权限组gid
	RGid string `gorm:"column:r_gid"`
	// 写权限组gid
	WGid string `gorm:"column:gid"`
	// 创建时间
	Ctime time.Time `gorm:"column:ctime;autoCreateTime"`
	// 修改时间
	Mtime time.Time `gorm:"column:mtime;autoUpdateTime"`
	// 删除时间
	Dtime gorm.DeletedAt `gorm:"column:dtime"`
	// 是否禁用（临时控制禁用，默认不禁用）
	Disabled bool `gomm:"column:disabled"`
	// 乐观锁
	Ver int `gorm:"column:ver"`
}
