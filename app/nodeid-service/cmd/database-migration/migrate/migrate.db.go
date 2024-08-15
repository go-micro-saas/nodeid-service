package dbmigrate

import (
	"context"
	dbv1_0_0 "github.com/go-micro-saas/nodeid-service/app/nodeid-service/cmd/database-migration/v1.0.0"
	setuputil "github.com/go-micro-saas/service-kit/setup"
	migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	stdlog "log"
)

// Run 运行迁移
func Run(launcherManager setuputil.LauncherManager, opts ...Option) {
	opt := &options{}
	for i := range opts {
		opts[i](opt)
	}

	// 关闭
	if opt.closeEngine {
		defer func() { _ = launcherManager.Close() }()
	}

	// 数据库链接
	//dbConn, err := setuputil.GetMysqlDBConn(launcher)
	dbConn, err := setuputil.GetPostgresDBConn(launcherManager)
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	// migrateHandler 迁移手柄
	var (
		ctx         = context.Background()
		migrateRepo = migrationpkg.NewMigrateRepo(dbConn)
	)

	// 初始化迁移记录
	if err = migrateRepo.InitializeSchema(ctx); err != nil {
		logpkg.Fatalw("migrateHandler.InitializeSchema failed", err)
	}

	// v1.0.0
	if err = dbv1_0_0.Upgrade(ctx, dbConn, migrateRepo); err != nil {
		logpkg.Fatalw("dbv1_0_0.Upgrade failed", err)
	}
}
