//go:build wireinject
// +build wireinject

package exportservices

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/biz"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/data"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/service/service"
	serverutil "github.com/go-micro-saas/service-kit/server"
	setuputil "github.com/go-micro-saas/service-kit/setup"
	"github.com/google/wire"
)

func initServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (*serverutil.Services, func(), error) {
	panic(wire.Build(
		// service
		setuputil.GetLogger,
		setuputil.GetPostgresDBConn,
		data.NewNodeIdData, data.NewNodeSerialData,
		biz.NewNodeIDBiz,
		service.NewNodeIDV1Service,
		// register services
		service.RegisterServices,
	))
	return &serverutil.Services{}, func() {}, nil
}
