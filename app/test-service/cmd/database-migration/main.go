package main

import (
	"flag"
	dbmigrate "github.com/go-micro-saas/nodeid-service/app/test-service/cmd/database-migration/migrate"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
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

// go run ./cmd/store-configuration/... -conf=./configs
func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	launcher, err := setuputil.NewLauncherManager(flagconf)
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	dbmigrate.Run(launcher, dbmigrate.WithCloseEngineHandler())
}
