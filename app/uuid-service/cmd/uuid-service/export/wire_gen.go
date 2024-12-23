// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-micro-saas/nodeid-service/api/nodeid-service/helper"
	"github.com/go-micro-saas/nodeid-service/api/uuid-service/helper"
	"github.com/go-micro-saas/nodeid-service/api/uuid-service/v1/services"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/cmd/nodeid-service/export"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/biz/biz"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/biz/repo"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/conf"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/service/dto"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/service/service"
	"github.com/ikaiguang/go-srv-kit/kit/id"
	"github.com/ikaiguang/go-srv-kit/service/cleanup"
	"github.com/ikaiguang/go-srv-kit/service/setup"
)

// Injectors from wire.go:

func exportNodeIDHelper(launcherManager setuputil.LauncherManager) (nodeidapi.NodeIDHelper, error) {
	srvNodeIDV1Server, err := serviceexporter.ExportNodeIDV1Service(launcherManager)
	if err != nil {
		return nil, err
	}
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	v := initNodeIDHelperOptions(logger)
	nodeIDHelper := nodeidapi.NewNodeIDHelper(srvNodeIDV1Server, v...)
	return nodeIDHelper, nil
}

func exportUuidHelper(launcherManager setuputil.LauncherManager) (uuidapi.UuidHelper, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	nodeIDHelper, err := exportNodeIDHelper(launcherManager)
	if err != nil {
		return nil, err
	}
	uuidHelper := uuidapi.NewUuidHelper(logger, nodeIDHelper)
	return uuidHelper, nil
}

func exportSnowflakeNode(launcherManager setuputil.LauncherManager) (idpkg.Snowflake, func(), error) {
	context := uuidapi.GetContext()
	uuidHelper, err := exportUuidHelper(launcherManager)
	if err != nil {
		return nil, nil, err
	}
	serviceConfig := conf.GetServiceConfig()
	getNodeIdReq := dto.ToPbGetNodeIdReq(serviceConfig)
	snowflake, cleanup, err := uuidapi.GetSnowflakeNode(context, uuidHelper, getNodeIdReq)
	if err != nil {
		return nil, nil, err
	}
	return snowflake, func() {
		cleanup()
	}, nil
}

func exportUuidBizRepo(launcherManager setuputil.LauncherManager) (bizrepos.UuidBizRepo, func(), error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, nil, err
	}
	snowflake, cleanup, err := exportSnowflakeNode(launcherManager)
	if err != nil {
		return nil, nil, err
	}
	uuidBizRepo := biz.NewUuidBiz(logger, snowflake)
	return uuidBizRepo, func() {
		cleanup()
	}, nil
}

func exportUuidV1Service(launcherManager setuputil.LauncherManager) (servicev1.SrvUuidV1Server, func(), error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, nil, err
	}
	uuidBizRepo, cleanup, err := exportUuidBizRepo(launcherManager)
	if err != nil {
		return nil, nil, err
	}
	srvUuidV1Server := service.NewUuidV1Service(logger, uuidBizRepo)
	return srvUuidV1Server, func() {
		cleanup()
	}, nil
}

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, func(), error) {
	srvUuidV1Server, cleanup, err := exportUuidV1Service(launcherManager)
	if err != nil {
		return nil, nil, err
	}
	cleanupManager, err := service.RegisterServices(hs, gs, srvUuidV1Server)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return cleanupManager, func() {
		cleanup()
	}, nil
}

// wire.go:

func initNodeIDHelperOptions(logger log.Logger) []nodeidapi.Option {
	return []nodeidapi.Option{nodeidapi.WithLogger(logger)}
}
