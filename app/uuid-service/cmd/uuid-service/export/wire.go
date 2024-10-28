/*
//go:build wireinject
// +build wireinject
*/

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	nodeidapi "github.com/go-micro-saas/nodeid-service/api/nodeid-service/helper"
	uuidapi "github.com/go-micro-saas/nodeid-service/api/uuid-service/helper"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/biz/biz"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/conf"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/service/dto"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/service/service"
	"github.com/google/wire"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (serverutil.ServiceInterface, func(), error) {
	panic(wire.Build(
		// service
		setuputil.GetLogger,
		nodeidapi.NewNodeIDHelper,
		uuidapi.GetContext, conf.GetServiceConfig, dto.ToPbGetNodeIdReq, uuidapi.GetSnowflakeNode,
		biz.NewUuidBiz,
		service.NewUuidV1Service,
		// register services
		service.RegisterServices,
	))
	return nil, nil, nil
}
