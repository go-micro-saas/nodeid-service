package daos

import (
	setuputil "github.com/go-micro-saas/service-kit/setup"
	"gorm.io/gorm"
	stdlog "log"
	"os"
	"testing"
)

var (
	dbConnection      *gorm.DB
	nodeIdHandler     *nodeIdRepo
	nodeSerialHandler *nodeSerialRepo
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

	dbConnection = db
	nodeIdHandler = NewNodeIdRepo(db).(*nodeIdRepo)
	nodeSerialHandler = NewNodeSerialRepo(db).(*nodeSerialRepo)

	os.Exit(m.Run())
}
