package dto

import (
	resourcev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/conf"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	NodeIDDto nodeIDDto
)

type nodeIDDto struct{}

func (s *nodeIDDto) ToBoNodeIDConfig(cfg *conf.ServiceConfig) *bo.NodeIDConfig {
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

func (s *nodeIDDto) ToBoGetNodeIdParam(req *resourcev1.GetNodeIdReq) *bo.GetNodeIdParam {
	res := &bo.GetNodeIdParam{
		InstanceId:   req.GetInstanceId(),
		InstanceName: req.GetInstanceName(),
		Metadata:     req.GetMetadata(),
	}

	return res
}

func (s *nodeIDDto) ToPbGetNodeIdRespData(cfg *bo.NodeIDConfig, dataModel *po.NodeId) *resourcev1.GetNodeIdRespData {
	res := &resourcev1.GetNodeIdRespData{
		Id:                dataModel.Id,
		InstanceId:        dataModel.InstanceId,
		InstanceName:      dataModel.InstanceName,
		NodeId:            dataModel.NodeId,
		Status:            dataModel.NodeIdStatus,
		ExpiredAt:         timestamppb.New(dataModel.ExpiredAt),
		HeartbeatInterval: durationpb.New(cfg.HeartbeatInterval),
	}
	return res
}

func (s *nodeIDDto) ToBoRenewalNodeIdParam(req *resourcev1.RenewalNodeIdReq) *bo.RenewalNodeIdParam {
	res := &bo.RenewalNodeIdParam{
		ID:         req.GetId(),
		InstanceId: req.GetInstanceId(),
		NodeID:     req.GetNodeId(),
	}

	return res
}

func (s *nodeIDDto) ToPbRenewalNodeIdRespData(dataModel *po.NodeId) *resourcev1.RenewalNodeIdRespData {
	res := &resourcev1.RenewalNodeIdRespData{
		Status:    dataModel.NodeIdStatus,
		ExpiredAt: timestamppb.New(dataModel.ExpiredAt),
	}
	return res
}
