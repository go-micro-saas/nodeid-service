package uuidapi

import (
	servicev1 "github.com/go-micro-saas/nodeid-service/api/uuid-service/v1/services"
	_ "github.com/gorilla/websocket"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
)

// GetAuthWhiteList 验证白名单
func GetAuthWhiteList() map[string]middlewareutil.TransportServiceKind {
	// 白名单
	whiteList := make(map[string]middlewareutil.TransportServiceKind)

	// 测试
	whiteList[servicev1.OperationSrvUuidV1NextID] = middlewareutil.TransportServiceKindALL

	return whiteList
}
