package dto

import (
	resourcev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	NodeIDDto nodeIDDto
)

type nodeIDDto struct{}

func (s *nodeIDDto) ToPbGetServiceInfoRespData(dataModel *bo.NodeIDConfig) *resourcev1.GetServiceInfoRespData {
	res := &resourcev1.GetServiceInfoRespData{
		MinNodeId:         uint64(dataModel.MinNodeID),
		MaxNodeId:         uint64(dataModel.MaxNodeID),
		IdleDuration:      durationpb.New(dataModel.IdleDuration),
		HeartbeatInterval: durationpb.New(dataModel.HeartbeatInterval),
	}
	if !dataModel.NodeEpoch.IsZero() {
		res.NodeEpoch = timestamppb.New(dataModel.NodeEpoch)
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
		InstanceId:        dataModel.InstanceId,
		NodeId:            dataModel.NodeId,
		Status:            dataModel.NodeIdStatus,
		ExpiredAt:         timestamppb.New(dataModel.ExpiredAt),
		HeartbeatInterval: durationpb.New(cfg.HeartbeatInterval),
		AccessToken:       dataModel.AccessToken,
	}
	if !cfg.NodeEpoch.IsZero() {
		res.NodeEpoch = timestamppb.New(cfg.NodeEpoch)
	}
	return res
}

func (s *nodeIDDto) ToBoRenewalNodeIdParam(req *resourcev1.RenewalNodeIdReq) *bo.RenewalNodeIdParam {
	res := &bo.RenewalNodeIdParam{
		InstanceId:  req.GetInstanceId(),
		NodeID:      req.GetNodeId(),
		AccessToken: req.GetAccessToken(),
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

func (s *nodeIDDto) ToPbRenewalNodeIdRespData2(dataModel *bo.RenewalNodeIDReply) *resourcev1.RenewalNodeIdRespData {
	res := &resourcev1.RenewalNodeIdRespData{
		Status:    dataModel.Status,
		ExpiredAt: timestamppb.New(dataModel.ExpiredAt),
	}
	return res
}

func (s *nodeIDDto) ToBoReleaseNodeIdParam(req *resourcev1.ReleaseNodeIdReq) *bo.ReleaseNodeIdParam {
	res := &bo.ReleaseNodeIdParam{
		InstanceId:  req.GetInstanceId(),
		NodeID:      req.GetNodeId(),
		AccessToken: req.GetAccessToken(),
	}

	return res
}

func (s *nodeIDDto) ToPbReleaseNodeIdRespData(dataModel *po.NodeId) *resourcev1.ReleaseNodeIdRespData {
	res := &resourcev1.ReleaseNodeIdRespData{
		Status:    dataModel.NodeIdStatus,
		ExpiredAt: timestamppb.New(dataModel.ExpiredAt),
	}
	return res
}
