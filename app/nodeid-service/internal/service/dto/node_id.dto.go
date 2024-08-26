package dto

import (
	resourcev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/conf"
	"google.golang.org/protobuf/types/known/durationpb"
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

func (s *nodeIDDto) ToPbGetServiceInfoRespData(dataModel *bo.NodeIDConfig) *resourcev1.GetServiceInfoRespData {
	res := &resourcev1.GetServiceInfoRespData{
		MinNodeId:         uint64(dataModel.MinNodeID),
		MaxNodeId:         uint64(dataModel.MaxNodeID),
		IdleDuration:      durationpb.New(dataModel.IdleDuration),
		HeartbeatInterval: durationpb.New(dataModel.HeartbeatInterval),
	}

	return res
}
