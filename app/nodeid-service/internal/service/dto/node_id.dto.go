package dto

import (
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/conf"
)

var (
	NodeIDDto nodeIDDto
)

type nodeIDDto struct{}

func (s *nodeIDDto) ToBoNodeIDConfig(cfg *conf.BusinessConfig) *bo.NodeIDConfig {
	res := &bo.NodeIDConfig{}
	if cfg.GetBusiness() == nil {
		return res
	}
	businessConfig := cfg.GetBusiness()
	res.MinNodeID = businessConfig.GetMinNodeId()
	res.MaxNodeID = businessConfig.GetMaxNodeId()
	res.IdleDuration = businessConfig.GetIdleDuration().AsDuration()
	res.HeartbeatInterval = businessConfig.GetHeartbeatInterval().AsDuration()
	return res
}
