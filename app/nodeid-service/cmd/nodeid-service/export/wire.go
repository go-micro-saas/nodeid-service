//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	servicev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/services"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/biz"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/conf"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/data"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/service/dto"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/service/service"
	"github.com/google/wire"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func exportNodeIDV1Service(launcherManager setuputil.LauncherManager) (servicev1.SrvNodeIDV1Server, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		setuputil.GetRecommendDBConn,
		//setuputil.GetRabbitmqConn,
		// data
		data.NewNodeIdData,
		data.NewNodeSerialData,
		//data.NewNodeEventHistoryRepo,
		// conf
		conf.GetServiceConfig,
		dto.ToBoNodeIDConfig,
		// biz
		biz.NewNodeIDBiz,
		// event
		//events.NewRenewNodeIDEventRepo,
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
