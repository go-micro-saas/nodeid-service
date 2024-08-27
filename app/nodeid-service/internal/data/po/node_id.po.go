package po

import (
	enumv1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/enums"
	datatypes "gorm.io/datatypes"
	time "time"
)

// NodeId ENGINE InnoDB CHARSET utf8mb4 COMMENT '节点'
type NodeId struct {
	Id               uint64                         `gorm:"column:id;primaryKey" json:"id"`                    // ID
	CreatedTime      time.Time                      `gorm:"column:created_time" json:"created_time"`           // 创建时间
	UpdatedTime      time.Time                      `gorm:"column:updated_time" json:"updated_time"`           // 更新时间
	InstanceName     string                         `gorm:"column:instance_name" json:"instance_name"`         // 实例名称
	InstanceId       string                         `gorm:"column:instance_id" json:"instance_id"`             // 实例ID
	NodeId           int64                          `gorm:"column:node_id" json:"node_id"`                     // 节点id
	NodeIdStatus     enumv1.NodeIDStatusEnum_Status `gorm:"column:node_id_status" json:"node_id_status"`       // 节点状态
	InstanceMetadata datatypes.JSON                 `gorm:"column:instance_metadata" json:"instance_metadata"` // 实例元数据
	ExpiredAt        time.Time                      `gorm:"column:expired_at" json:"expired_at"`               // 失效时间
}
