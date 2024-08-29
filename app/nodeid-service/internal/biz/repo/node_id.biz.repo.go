package bizrepos

import (
	"context"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
)

type NodeIdBizRepo interface {
	GetConfig() *bo.NodeIDConfig
	GetNodeId(ctx context.Context, param *bo.GetNodeIdParam) (*po.NodeId, error)
	GenerateNextID(serialModel *po.NodeSerial, param *bo.GetNodeIdParam) (*po.NodeId, error)
	RenewalNodeId(ctx context.Context, param *bo.RenewalNodeIdParam) (*po.NodeId, error)
}
