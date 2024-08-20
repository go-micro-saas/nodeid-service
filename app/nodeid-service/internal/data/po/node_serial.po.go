package po

import (
	time "time"
)

// NodeSerial ENGINE InnoDB CHARSET utf8mb4 COMMENT '节点系列号'
type NodeSerial struct {
	Id            int64     `gorm:"column:id;primaryKey" json:"id"`                // ID
	CreatedTime   time.Time `gorm:"column:created_time" json:"created_time"`       // 创建时间
	UpdatedTime   time.Time `gorm:"column:updated_time" json:"updated_time"`       // 更新时间
	InstanceId    string    `gorm:"column:instance_id" json:"instance_id"`         // 实例ID
	CurrentNodeId int64     `gorm:"column:current_node_id" json:"current_node_id"` // 当前节点id
}
