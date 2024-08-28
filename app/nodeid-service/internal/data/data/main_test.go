package data

import (
	setuputil "github.com/go-micro-saas/service-kit/setup"
	"gorm.io/gorm"
	stdlog "log"
	"os"
	"testing"
)

var (
	dbConnection      *gorm.DB
	nodeIdHandler     *nodeIdData
	nodeSerialHandler *nodeSerialData
)

func TestMain(m *testing.M) {
	configPath := "./../../../configs"
	launcher, err := setuputil.NewLauncherManager(configPath)
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	db, err := launcher.GetPostgresDBConn()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	logger, err := launcher.GetLogger()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	dbConnection = db
	nodeIdHandler = NewNodeIdData(logger, db).(*nodeIdData)
	nodeSerialHandler = NewNodeSerialData(logger, db).(*nodeSerialData)

	os.Exit(m.Run())
}
