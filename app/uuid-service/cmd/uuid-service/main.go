package main

import (
	"flag"
	nodeidserviceexporter "github.com/go-micro-saas/nodeid-service/app/nodeid-service/cmd/nodeid-service/export"
	serviceexporter "github.com/go-micro-saas/nodeid-service/app/uuid-service/cmd/uuid-service/export"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
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

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	configOpts := serviceexporter.ExportServiceConfig()
	whitelist := serviceexporter.ExportAuthWhitelist()
	services := []serverutil.ServiceExporter{serviceexporter.ExportServices}

	// nodeid
	configOpts = append(configOpts, nodeidserviceexporter.ExportServiceConfig()...)

	app, cleanup, err := serverutil.AllInOneServer(flagconf, configOpts, services, whitelist)
	if err != nil {
		stdlog.Fatalf("==> runservices.GetServerApp failed: %+v\n", err)
	}
	serverutil.RunServer(app, cleanup)
}
