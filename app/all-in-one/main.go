package main

import (
	"flag"
	serviceexporter "github.com/go-micro-saas/nodeid-service/app/nodeid-service/cmd/nodeid-service/export"
	uuidserviceexporter "github.com/go-micro-saas/nodeid-service/app/uuid-service/cmd/uuid-service/export"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
	pingserviceexporter "github.com/ikaiguang/go-srv-kit/testdata/ping-service/cmd/ping-service/export"
	stdlog "log"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string

	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

// go run ./testdata/all-in-one/main.go -conf=./app/nodeid-service/configs
func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	var (
		configOpts []configutil.Option
		whitelist  []map[string]middlewareutil.TransportServiceKind
		services   []serverutil.ServiceExporter
	)

	// nodeid-service
	configOpts = append(configOpts, serviceexporter.ExportServiceConfig()...)
	whitelist = append(whitelist, serviceexporter.ExportAuthWhitelist()...)
	services = append(services, serviceexporter.ExportServices)

	// ping-service
	configOpts = append(configOpts, pingserviceexporter.ExportServiceConfig()...)
	whitelist = append(whitelist, pingserviceexporter.ExportAuthWhitelist()...)
	services = append(services, pingserviceexporter.ExportServices)

	// uuid-service
	configOpts = append(configOpts, uuidserviceexporter.ExportServiceConfig()...)
	whitelist = append(whitelist, uuidserviceexporter.ExportAuthWhitelist()...)
	services = append(services, uuidserviceexporter.ExportServices)

	app, cleanup, err := serverutil.AllInOneServer(flagconf, configOpts, services, whitelist)
	if err != nil {
		stdlog.Fatalf("==> serverutil.AllInOneServer failed: %+v\n", err)
	}
	serverutil.RunServer(app, cleanup)
}
