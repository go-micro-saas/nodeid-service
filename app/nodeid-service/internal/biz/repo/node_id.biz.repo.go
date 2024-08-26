package bizrepos

import "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"

type NodeIdBizRepo interface {
	GetConfig() *bo.NodeIDConfig
}
