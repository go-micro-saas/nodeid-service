//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	servicev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/services"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/biz"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/repo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/conf"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/data"
	datarepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/repo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/service/dto"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/service/service"
	"github.com/google/wire"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func exportNodeIdData(launcherManager setuputil.LauncherManager) (datarepos.NodeIdDataRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		setuputil.GetRecommendDBConn, data.NewNodeIdData,
	))
}

func exportNodeSerialData(launcherManager setuputil.LauncherManager) (datarepos.NodeSerialDataRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		setuputil.GetRecommendDBConn, data.NewNodeSerialData,
	))
}

func exportNodeIdBizRepo(launcherManager setuputil.LauncherManager) (bizrepos.NodeIdBizRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		exportNodeIdData, exportNodeSerialData,
		// biz
		conf.GetServiceConfig, dto.ToBoNodeIDConfig, biz.NewNodeIDBiz,
	))
}

func exportNodeIDV1Service(launcherManager setuputil.LauncherManager) (servicev1.SrvNodeIDV1Server, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// biz
		exportNodeIdBizRepo,
		// service
		service.NewNodeIDV1Service,
	))
}

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	panic(wire.Build(
		// service
		exportNodeIDV1Service,
		// register services
		service.RegisterServices,
	))
}
