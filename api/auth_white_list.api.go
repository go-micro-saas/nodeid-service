package nodeidapi

import (
	servicev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/services"
	middlewareutil "github.com/go-micro-saas/service-kit/middleware"
)

// GetAuthWhiteList 验证白名单
func GetAuthWhiteList() map[string]middlewareutil.TransportServiceKind {
	// 白名单
	whiteList := make(map[string]middlewareutil.TransportServiceKind)

	// 测试
	whiteList[servicev1.OperationSrvNodeIDV1Ping] = middlewareutil.TransportServiceKindALL

	return whiteList
}
