package dbv1_0_0_nodeid

import (
	"context"
	nodeeventschemas "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/schema/node_event_history"
	nodeidschemas "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/schema/node_id"
	nodeserialschemas "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/schema/node_serial"
	migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"gorm.io/gorm"
)

// Migrate 数据库迁移
type Migrate struct {
	dbConn      *gorm.DB
	migrateRepo migrationpkg.MigrateRepo
}

// NewMigrateHandler 处理手柄
func NewMigrateHandler(dbConn *gorm.DB, migrateRepo migrationpkg.MigrateRepo) *Migrate {
	return &Migrate{
		dbConn:      dbConn,
		migrateRepo: migrateRepo,
	}
}

// Upgrade ...
func (s *Migrate) Upgrade(ctx context.Context) error {
	var (
		mr       migrationpkg.MigrationInterface
		migrator = s.dbConn.WithContext(ctx).Migrator()
	)

	// 创建表
	mr = nodeidschemas.NodeIdSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 创建索引
	mr = nodeidschemas.NodeIdSchema.CreateUniqueIndexForInstanceIDAndNodeID(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 创建表
	mr = nodeserialschemas.NodeSerialSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 创建表
	mr = nodeeventschemas.NodeEventHistorySchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 添加字段
	mr = nodeidschemas.NodeIdSchema.AddColumnAccessToken(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	return nil
}
