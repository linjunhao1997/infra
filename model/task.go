package model

import "time"

var (
	// 待处理
	TaskStatePending = "Pending"
	// 处理中
	TaskStateHanding = "Handing"
	// 挂起
	TaskStateSuspended = "Suspended"
	// 转发（转发生成新任务单）
	TaskStateForwarded = "Forwarded"
	// 已解决
	TaskStateResolved = "Resolved"
)

// 任务型通用模型
type Task struct {
	// 主键id
	Id int64 `gorm:"column:id;primaryKey"`
	// 事件id
	EventId int64 `gorm:"column:event_id"`
	Event   Event
	// 父任务id
	SupId int64 `gorm:"column:sup_id"`
	// 派发人
	DispatcherUid string `gorm:"column:dispatcher_uid"`
	// 所属人Uid
	OwnerUid string `gorm:"column:owner_uid"`
	// 修改人uid
	ModifierUid string `gorm:"column:modifier_uid"`
	// 读权限组gid
	RGid string `gorm:"column:r_gid"`
	// 创建时间
	Ctime time.Time `gorm:"column:ctime"`
	// 修改时间
	Mtime time.Time `gorm:"column:mtime"`
	// 状态
	State string `gomm:"column:state"`
}

type Tasks []*Task
