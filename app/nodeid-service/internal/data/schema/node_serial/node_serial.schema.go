// Package schemas
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package schemas

import (
	migrationuitl "github.com/ikaiguang/go-srv-kit/data/migration"
	datatypes "gorm.io/datatypes"
	gorm "gorm.io/gorm"
	time "time"
)

var _ = time.Time{}

var _ = datatypes.JSON{}

// NodeSerialSchema NodeSerial
var NodeSerialSchema NodeSerial

// NewNodeSerial new schema
func NewNodeSerial() *NodeSerial {
	return &NodeSerial{}
}

const (
	TableName = "nid_node_serial"

	FieldId            = "id"
	FieldCreatedTime   = "created_time"
	FieldUpdatedTime   = "updated_time"
	FieldInstanceId    = "instance_id"
	FieldCurrentNodeId = "current_node_id"
)

// NodeSerial ENGINE InnoDB CHARSET utf8mb4 COMMENT '节点系列号'
type NodeSerial struct {
	Id            int64     `gorm:"column:id;primaryKey;type:int;autoIncrement;not null;default:;comment:ID" json:"id"`
	CreatedTime   time.Time `gorm:"column:created_time;type:time;not null;default:current_timestamp();comment:创建时间" json:"created_time"`
	UpdatedTime   time.Time `gorm:"column:updated_time;type:time;not null;default:current_timestamp();comment:更新时间" json:"updated_time"`
	InstanceId    string    `gorm:"column:instance_id;unique;type:string;size:255;not null;default:'';comment:实例ID" json:"instance_id"`
	CurrentNodeId int64     `gorm:"column:current_node_id;type:int;not null;default:0;comment:当前节点id" json:"current_node_id"`
}

// TableName table name
func (s *NodeSerial) TableName() string {
	return TableName
}

// CreateTableMigrator create table migrator
func (s *NodeSerial) CreateTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationInterface {
	return migrationuitl.NewCreateTable(migrator, migrationuitl.Version, s)
}

// DropTableMigrator create table migrator
func (s *NodeSerial) DropTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationInterface {
	return migrationuitl.NewDropTable(migrator, migrationuitl.Version, s)
}