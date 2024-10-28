/*
//go:build wireinject
// +build wireinject
*/

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/conf"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/service/service"
	"github.com/google/wire"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (serverutil.ServiceInterface, error) {
	panic(wire.Build(
		// service
		setuputil.GetLogger,
		setuputil.GetRecommendDBConn,
		data.NewNodeIdData, data.NewNodeSerialData,
		conf.GetServiceConfig, dto.ToBoNodeIDConfig,
		biz.NewNodeIDBiz,
		service.NewNodeIDV1Service,
		// register services
		service.RegisterServices,
	))
	return nil, nil
}
