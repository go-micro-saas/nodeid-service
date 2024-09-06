package nodeidapi

import (
	servicev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/services"
	middlewareutil "github.com/go-micro-saas/service-kit/middleware"
	_ "github.com/gorilla/websocket"
)

// GetAuthWhiteList 验证白名单
func GetAuthWhiteList() map[string]middlewareutil.TransportServiceKind {
	// 白名单
	whiteList := make(map[string]middlewareutil.TransportServiceKind)

	// 测试
	whiteList[servicev1.OperationSrvNodeIDV1Ping] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvNodeIDV1GetServiceInfo] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvNodeIDV1GetNodeId] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvNodeIDV1RenewalNodeId] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvNodeIDV1ReleaseNodeId] = middlewareutil.TransportServiceKindALL

	return whiteList
}
