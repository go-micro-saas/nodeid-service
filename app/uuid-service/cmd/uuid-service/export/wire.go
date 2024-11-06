//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	nodeidapi "github.com/go-micro-saas/nodeid-service/api/nodeid-service/helper"
	uuidapi "github.com/go-micro-saas/nodeid-service/api/uuid-service/helper"
	servicev1 "github.com/go-micro-saas/nodeid-service/api/uuid-service/v1/services"
	serviceexporter "github.com/go-micro-saas/nodeid-service/app/nodeid-service/cmd/nodeid-service/export"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/biz/biz"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/biz/repo"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/conf"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/service/dto"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/service/service"
	"github.com/google/wire"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func initNodeIDHelperOptions(logger log.Logger) []nodeidapi.Option {
	return []nodeidapi.Option{nodeidapi.WithLogger(logger)}
}

func exportNodeIDHelper(launcherManager setuputil.LauncherManager) (nodeidapi.NodeIDHelper, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// node service
		serviceexporter.ExportNodeIDV1Service,
		// nodeid helper
		initNodeIDHelperOptions, nodeidapi.NewNodeIDHelper,
	))
	return nil, nil
}

func exportUuidHelper(launcherManager setuputil.LauncherManager) (uuidapi.UuidHelper, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// nodeid helper
		exportNodeIDHelper,
		// uuid helper
		uuidapi.NewUuidHelper,
	))
	return nil, nil
}

func exportSnowflakeNode(launcherManager setuputil.LauncherManager) (idpkg.Snowflake, func(), error) {
	panic(wire.Build(
		// uuid helper
		exportUuidHelper,
		// snowflake
		uuidapi.GetContext, conf.GetServiceConfig, dto.ToPbGetNodeIdReq, uuidapi.GetSnowflakeNode,
	))
	return nil, nil, nil
}

func exportUuidBizRepo(launcherManager setuputil.LauncherManager) (bizrepos.UuidBizRepo, func(), error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// snowflake
		exportSnowflakeNode,
		// biz
		biz.NewUuidBiz,
	))
	return nil, nil, nil
}

func exportUuidV1Service(launcherManager setuputil.LauncherManager) (servicev1.SrvUuidV1Server, func(), error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// biz
		exportUuidBizRepo,
		// service
		service.NewUuidV1Service,
	))
	return nil, nil, nil
}

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, func(), error) {
	panic(wire.Build(
		// service
		exportUuidV1Service,
		// register services
		service.RegisterServices,
	))
	return nil, nil, nil
}
