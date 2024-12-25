//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	nodeidapi "github.com/go-micro-saas/nodeid-service/api/nodeid-service/helper"
	uuidapi "github.com/go-micro-saas/nodeid-service/api/uuid-service/helper"
	serviceexporter "github.com/go-micro-saas/nodeid-service/app/nodeid-service/cmd/nodeid-service/export"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/biz/biz"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/conf"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/service/dto"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/service/service"
	"github.com/google/wire"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, func(), error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// conf
		uuidapi.GetContext,
		conf.GetServiceConfig,
		dto.ToPbGetNodeIdReq,
		// node service
		serviceexporter.ExportNodeIDV1Service,
		// nodeid helper
		nodeidapi.DefaultOptions, nodeidapi.NewNodeIDHelper,
		// uuid helper
		uuidapi.NewUuidHelper,
		// snowflake
		uuidapi.GetSnowflakeNode,
		// biz
		biz.NewUuidBiz,
		// service
		service.NewUuidV1Service,
		// register services
		service.RegisterServices,
	))
	return nil, nil, nil
}
