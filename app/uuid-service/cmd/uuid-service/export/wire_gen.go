// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-micro-saas/nodeid-service/api/nodeid-service/helper"
	"github.com/go-micro-saas/nodeid-service/api/uuid-service/helper"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/cmd/nodeid-service/export"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/biz/biz"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/conf"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/service/dto"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/service/service"
	"github.com/ikaiguang/go-srv-kit/service/cleanup"
	"github.com/ikaiguang/go-srv-kit/service/setup"
)

// Injectors from wire.go:

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, func(), error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, nil, err
	}
	context := uuidapi.GetContext()
	srvNodeIDV1Server, err := serviceexporter.ExportNodeIDV1Service(launcherManager)
	if err != nil {
		return nil, nil, err
	}
	v := nodeidapi.DefaultOptions(logger)
	nodeIDHelper := nodeidapi.NewNodeIDHelper(srvNodeIDV1Server, v...)
	uuidHelper := uuidapi.NewUuidHelper(logger, nodeIDHelper)
	serviceConfig := conf.GetServiceConfig()
	getNodeIdReq := dto.ToPbGetNodeIdReq(serviceConfig)
	snowflake, cleanup, err := uuidapi.GetSnowflakeNode(context, uuidHelper, getNodeIdReq)
	if err != nil {
		return nil, nil, err
	}
	uuidBizRepo := biz.NewUuidBiz(logger, snowflake)
	srvUuidV1Server := service.NewUuidV1Service(logger, uuidBizRepo)
	cleanupManager, err := service.RegisterServices(hs, gs, srvUuidV1Server)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return cleanupManager, func() {
		cleanup()
	}, nil
}
