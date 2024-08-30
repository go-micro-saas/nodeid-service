package serviceexporter

import (
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/conf"
	configutil "github.com/go-micro-saas/service-kit/config"
	middlewareutil "github.com/go-micro-saas/service-kit/middleware"
	serverutil "github.com/go-micro-saas/service-kit/server"
	setuputil "github.com/go-micro-saas/service-kit/setup"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/api"
)

func ExportServiceConfig() []configutil.Option {
	return conf.LoadServiceConfig()
}

func ExportAuthWhitelist() []map[string]middlewareutil.TransportServiceKind {
	return []map[string]middlewareutil.TransportServiceKind{
		api.GetAuthWhiteList(),
	}
}

func ExportServices(launcherManager setuputil.LauncherManager, serverManager serverutil.ServerManager) (serverutil.ServiceInterface, error) {
	hs, err := serverManager.GetHTTPServer()
	if err != nil {
		return nil, err
	}
	gs, err := serverManager.GetGRPCServer()
	if err != nil {
		return nil, err
	}
	return exportServices(launcherManager, hs, gs)
}
