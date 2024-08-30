package exportservices

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/conf"
	configutil "github.com/go-micro-saas/service-kit/config"
	serverutil "github.com/go-micro-saas/service-kit/server"
	setuputil "github.com/go-micro-saas/service-kit/setup"
)

func ExportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (*serverutil.Services, func(), error) {
	return initServices(launcherManager, hs, gs)
}

func LoadServiceConfig() configutil.Option {
	return conf.LoadServiceConfig()
}
