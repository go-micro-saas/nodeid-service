// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/services"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/biz"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/repo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/conf"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/data"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/repo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/service/dto"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/service/service"
	"github.com/ikaiguang/go-srv-kit/service/cleanup"
	"github.com/ikaiguang/go-srv-kit/service/setup"
)

// Injectors from wire.go:

func exportNodeIdData(launcherManager setuputil.LauncherManager) (datarepos.NodeIdDataRepo, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	db, err := setuputil.GetRecommendDBConn(launcherManager)
	if err != nil {
		return nil, err
	}
	nodeIdDataRepo := data.NewNodeIdData(logger, db)
	return nodeIdDataRepo, nil
}

func exportNodeSerialData(launcherManager setuputil.LauncherManager) (datarepos.NodeSerialDataRepo, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	db, err := setuputil.GetRecommendDBConn(launcherManager)
	if err != nil {
		return nil, err
	}
	nodeSerialDataRepo := data.NewNodeSerialData(logger, db)
	return nodeSerialDataRepo, nil
}

func exportNodeIdBizRepo(launcherManager setuputil.LauncherManager) (bizrepos.NodeIdBizRepo, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	serviceConfig := conf.GetServiceConfig()
	nodeIDConfig := dto.ToBoNodeIDConfig(serviceConfig)
	nodeIdDataRepo, err := exportNodeIdData(launcherManager)
	if err != nil {
		return nil, err
	}
	nodeSerialDataRepo, err := exportNodeSerialData(launcherManager)
	if err != nil {
		return nil, err
	}
	nodeIdBizRepo := biz.NewNodeIDBiz(logger, nodeIDConfig, nodeIdDataRepo, nodeSerialDataRepo)
	return nodeIdBizRepo, nil
}

func exportNodeIDV1Service(launcherManager setuputil.LauncherManager) (servicev1.SrvNodeIDV1Server, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	nodeIdBizRepo, err := exportNodeIdBizRepo(launcherManager)
	if err != nil {
		return nil, err
	}
	srvNodeIDV1Server := service.NewNodeIDV1Service(logger, nodeIdBizRepo)
	return srvNodeIDV1Server, nil
}

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	srvNodeIDV1Server, err := exportNodeIDV1Service(launcherManager)
	if err != nil {
		return nil, err
	}
	cleanupManager, err := service.RegisterServices(hs, gs, srvNodeIDV1Server)
	if err != nil {
		return nil, err
	}
	return cleanupManager, nil
}
