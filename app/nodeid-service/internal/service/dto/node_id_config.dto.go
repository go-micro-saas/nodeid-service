package dto

import (
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/conf"
)

func ToBoNodeIDConfig(cfg *conf.ServiceConfig) *bo.NodeIDConfig {
	res := &bo.NodeIDConfig{}
	if cfg.GetNodeidService() == nil {
		return res
	}
	serviceConfig := cfg.GetNodeidService()
	res.MinNodeID = serviceConfig.GetMinNodeId()
	res.MaxNodeID = serviceConfig.GetMaxNodeId()
	if serviceConfig.GetNodeEpoch() != nil {
		res.NodeEpoch = serviceConfig.GetNodeEpoch().AsTime()
	}
	res.IdleDuration = serviceConfig.GetIdleDuration().AsDuration()
	res.HeartbeatInterval = serviceConfig.GetHeartbeatInterval().AsDuration()
	return res
}
