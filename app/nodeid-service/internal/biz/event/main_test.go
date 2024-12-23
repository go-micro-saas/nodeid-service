package events

import (
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/biz"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/data"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	stdlog "log"
	"os"
	"testing"
)

var (
	handler *renewNodeIDEvent
)

func mockConfig() *bo.NodeIDConfig {
	res := &bo.NodeIDConfig{
		MinNodeID:         0,
		MaxNodeID:         0,
		IdleDuration:      0,
		HeartbeatInterval: 0,
	}
	res.Initialization()
	return res
}

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

	mqConn, err := launcher.GetRabbitmqConn()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	nodeIDData := data.NewNodeIdData(logger, db)
	nodeSerialIDData := data.NewNodeSerialData(logger, db)
	nodeidBizRepo := biz.NewNodeIDBiz(logger, mockConfig(), nodeIDData, nodeSerialIDData)

	handler = NewRenewNodeIDEventRepo(logger, mqConn, nodeidBizRepo).(*renewNodeIDEvent)

	os.Exit(m.Run())
}
