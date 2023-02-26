package model

import "time"

// 任务状态
var (
	// 已创建，保存触发
	EventStateCreated = "Created"
	// 已提交，提交触发
	EventStateCommited = "Commited"
	// 撤回, 撤回触发
	EventStateWithdrew = "Withdrew"
	// 关闭，关闭触发
	EventStateClosed = "Closed"

	// 处理中，有未完成的任务
	EventStateHanding = "Handing"
	// 已结束，任务全部完成，有可能产生下一类型事件
	EventStateFinished = "Finished"
)

// 事件型通用模型, 如问题单，工作单
// 每个事件可产生多个子事件
// 每个事件都需要分配任务，除了直接操作事件的状态外，事件状态的更新由任务的处理情况决定
type Event struct {
	// 主键id
	Id int64 `gorm:"column:id;primaryKey"`
	// 标题
	Title string `gorm:"column:title"`
	// 描述 （此字段可拓展成多个业务相关的字段）
	Detail string `gorm:"column:detail"`
	// 编号
	Sn string `gorm:"column:sn"`
	// 状态
	State string `gorm:"column:state"`
	// 上级事件id
	SupId int64 `gorm:"column:sup_id"`
	// 上级事件
	Sup *Event
	// 创建人uid（事件的创建人，有可能是机器人创建的事件）
	CreatorUid string `gorm:"column:CreatorUid"`
	// 修改人uid
	ModifierUid string `gorm:"column:modifier_uid"`
	// 读权限组gid
	RGid string `gorm:"column:r_gid"`
	// 写权限组gid（可由组里面的人安排任务）
	WGid string `gorm:"column:gid"`
	// 创建时间
	Ctime time.Time `gorm:"column:ctime;autoCreateTime"`
	// 修改时间
	Mtime time.Time `gorm:"column:mtime;autoUpdateTime"`
	// 开始时间
	Stime time.Time `gorm:"column:stime"`
	// 结束时间
	Etime time.Time `gorm:"column:etime"`
	// 事件安排的任务
	Tasks Tasks
}
