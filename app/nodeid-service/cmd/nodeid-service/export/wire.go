//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	servicev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/services"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/biz"
	events "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/event"
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
	return nil, nil
}

func exportNodeSerialData(launcherManager setuputil.LauncherManager) (datarepos.NodeSerialDataRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		setuputil.GetRecommendDBConn, data.NewNodeSerialData,
	))
	return nil, nil
}

func exportNodeIdBizRepo(launcherManager setuputil.LauncherManager) (bizrepos.NodeIdBizRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		exportNodeIdData, exportNodeSerialData,
		// biz
		conf.GetServiceConfig, dto.ToBoNodeIDConfig, biz.NewNodeIDBiz,
	))
	return nil, nil
}

func exportRenewNodeIDEventRepo(launcherManager setuputil.LauncherManager) (bizrepos.RenewNodeIDEventRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		setuputil.GetRabbitmqConn,
		// biz
		exportNodeIdBizRepo,
		// event
		events.NewRenewNodeIDEventRepo,
	))
	return nil, nil
}

func exportNodeIDV1Service(launcherManager setuputil.LauncherManager) (servicev1.SrvNodeIDV1Server, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// biz
		exportNodeIdBizRepo,
		// event
		exportRenewNodeIDEventRepo,
		// service
		service.NewNodeIDV1Service,
	))
	return nil, nil
}

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	panic(wire.Build(
		// service
		exportNodeIDV1Service,
		// register services
		service.RegisterServices,
	))
	return nil, nil
}
