// Package schemas
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package schemas

import (
	migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
	datatypes "gorm.io/datatypes"
	gorm "gorm.io/gorm"
	time "time"
)

var _ = time.Time{}

var _ = datatypes.JSON{}

// NodeIdSchema NodeId
var NodeIdSchema NodeId

// NewNodeId new schema
func NewNodeId() *NodeId {
	return &NodeId{}
}

const (
	TableName = "nid_node_id"

	FieldId               = "id"
	FieldCreatedTime      = "created_time"
	FieldUpdatedTime      = "updated_time"
	FieldInstanceName     = "instance_name"
	FieldInstanceId       = "instance_id"
	FieldNodeId           = "node_id"
	FieldNodeIdStatus     = "node_id_status"
	FieldInstanceMetadata = "instance_metadata"
	FieldExpiredAt        = "expired_at"
)

// NodeId ENGINE InnoDB CHARSET utf8mb4 COMMENT '节点'
type NodeId struct {
	Id               uint64         `gorm:"column:id;primaryKey;type:uint;autoIncrement;not null;default:;comment:ID" json:"id"`
	CreatedTime      time.Time      `gorm:"column:created_time;type:time;not null;comment:创建时间" json:"created_time"`
	UpdatedTime      time.Time      `gorm:"column:updated_time;type:time;not null;comment:更新时间" json:"updated_time"`
	InstanceName     string         `gorm:"column:instance_name;type:string;size:255;not null;default:'';comment:实例名称" json:"instance_name"`
	InstanceId       string         `gorm:"column:instance_id;type:string;size:255;not null;default:'';comment:实例ID" json:"instance_id"`
	NodeId           int64          `gorm:"column:node_id;type:int;not null;default:0;comment:节点id" json:"node_id"`
	NodeIdStatus     int32          `gorm:"column:node_id_status;type:int;not null;default:0;comment:节点状态" json:"node_id_status"`
	InstanceMetadata datatypes.JSON `gorm:"column:instance_metadata;type:json;not null;comment:实例元数据" json:"instance_metadata"`
	ExpiredAt        time.Time      `gorm:"column:expired_at;index;type:time;not null;comment:失效时间" json:"expired_at"`
	AccessToken      string         `gorm:"column:access_token;type:string;size:255;not null;default:'';comment:令牌；用于续订和释放ID" json:"access_token"`
}

// TableName table name
func (s *NodeId) TableName() string {
	return TableName
}

// CreateTableMigrator create table migrator
func (s *NodeId) CreateTableMigrator(migrator gorm.Migrator) migrationpkg.MigrationInterface {
	return migrationpkg.NewCreateTable(migrator, migrationpkg.Version, s)
}

// DropTableMigrator create table migrator
func (s *NodeId) DropTableMigrator(migrator gorm.Migrator) migrationpkg.MigrationInterface {
	return migrationpkg.NewDropTable(migrator, migrationpkg.Version, s)
}

// InstanceNodeUniqueIndex ...
type InstanceNodeUniqueIndex struct {
	InstanceId string `gorm:"column:instance_id;uniqueIndex:idx_node_id_instance_id_node_id;type:string;size:255;not null;default:'';comment:实例ID" json:"instance_id"`
	NodeId     int64  `gorm:"column:node_id;;uniqueIndex:idx_node_id_instance_id_node_id;type:int;not null;default:0;comment:节点id" json:"node_id"`
}

// TableName table name
func (s *InstanceNodeUniqueIndex) TableName() string {
	return NodeIdSchema.TableName()
}

func (s *InstanceNodeUniqueIndex) IndexName() string {
	return "idx_node_id_instance_id_node_id"
}

// CreateUniqueIndexForInstanceIDAndNodeID 创建唯一索引
func (s *NodeId) CreateUniqueIndexForInstanceIDAndNodeID(migrator gorm.Migrator) migrationpkg.MigrationInterface {
	var (
		dataModel           = &InstanceNodeUniqueIndex{}
		indexName           = dataModel.IndexName()
		migrationVersion    = migrationpkg.Version
		migrationIdentifier = migrationVersion + ":" + s.TableName() + ":create_unique_index:" + indexName
	)
	migrationUp := func() error {
		if migrator.HasIndex(dataModel, indexName) {
			return nil
		}
		return migrator.CreateIndex(dataModel, indexName)
	}
	migrationDown := func() error {
		if !migrator.HasIndex(dataModel, indexName) {
			return nil
		}
		return migrator.DropIndex(dataModel, indexName)
	}

	return migrationpkg.NewAnyMigrator(
		migrationVersion,
		migrationIdentifier,
		migrationUp,
		migrationDown,
	)
}

// AddColumnAccessToken 添加字段
func (s *NodeId) AddColumnAccessToken(migrator gorm.Migrator) migrationpkg.MigrationInterface {
	var (
		dataModel           = &NodeId{}
		columnName          = "access_token"
		migrationVersion    = migrationpkg.Version
		migrationIdentifier = migrationVersion + ":" + s.TableName() + ":add_column:" + columnName
	)
	migrationUp := func() error {
		if migrator.HasColumn(dataModel, columnName) {
			return nil
		}
		return migrator.AddColumn(dataModel, columnName)
	}
	migrationDown := func() error {
		if !migrator.HasColumn(dataModel, columnName) {
			return nil
		}
		return migrator.DropColumn(dataModel, columnName)
	}

	return migrationpkg.NewAnyMigrator(
		migrationVersion,
		migrationIdentifier,
		migrationUp,
		migrationDown,
	)
}

// TableSQL table SQL
func (s *NodeId) TableSQL() string {
	return `
CREATE TABLE if NOT EXISTS nid_node_id (
	id bigint unsigned NOT NULL auto_increment comment 'ID',
	created_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP() comment '创建时间',
	updated_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP() comment '更新时间',
	instance_name VARCHAR(255) NOT NULL DEFAULT '' comment '实例名称',
	instance_id VARCHAR(255) NOT NULL DEFAULT '' comment '实例ID',
	node_id bigint NOT NULL DEFAULT 0 comment '节点id',
	node_id_status mediumint NOT NULL DEFAULT 0 comment '节点状态',
	instance_metadata json NOT NULL comment '实例元数据',
	expired_at datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP() comment '失效时间',
	PRIMARY KEY (id),
	UNIQUE KEY (instance_id, node_id),
	KEY (expired_at)
) ENGINE InnoDB,
  CHARSET utf8mb4,
  COMMENT '节点'
`
}
