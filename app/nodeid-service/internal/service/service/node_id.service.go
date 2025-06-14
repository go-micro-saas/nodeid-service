package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	enumv1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/enums"
	resourcev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	servicev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/services"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/repo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/service/dto"
	"time"
)

type nodeIDV1Service struct {
	servicev1.UnimplementedSrvNodeIDV1Server

	log              *log.Helper
	nodeIDBiz        bizrepos.NodeIdBizRepo
	renewNodeIDEvent bizrepos.RenewNodeIDEventRepo
}

func NewNodeIDV1Service(
	logger log.Logger,
	nodeIDBiz bizrepos.NodeIdBizRepo,
	// renewNodeIDEvent bizrepos.RenewNodeIDEventRepo,
) servicev1.SrvNodeIDV1Server {
	logHelper := log.NewHelper(log.With(logger, "module", "nodeid-service/service/service"))
	return &nodeIDV1Service{
		log:       logHelper,
		nodeIDBiz: nodeIDBiz,
		//renewNodeIDEvent: renewNodeIDEvent,
	}
}

// Ping ping
func (s *nodeIDV1Service) Ping(ctx context.Context, req *resourcev1.PingReq) (*resourcev1.PingResp, error) {
	return &resourcev1.PingResp{Data: &resourcev1.PingRespData{
		Message: "request message: " + req.Message,
	}}, nil
}

// GetServiceInfo 获取服务信息
func (s *nodeIDV1Service) GetServiceInfo(ctx context.Context, req *resourcev1.GetServiceInfoReq) (*resourcev1.GetServiceInfoResp, error) {
	cfg := s.nodeIDBiz.GetConfig()

	return &resourcev1.GetServiceInfoResp{
		Data: dto.NodeIDDto.ToPbGetServiceInfoRespData(cfg),
	}, nil
}

// GetNodeId 获取节点id
func (s *nodeIDV1Service) GetNodeId(ctx context.Context, req *resourcev1.GetNodeIdReq) (*resourcev1.GetNodeIdResp, error) {
	param := dto.NodeIDDto.ToBoGetNodeIdParam(req)
	dataModel, err := s.nodeIDBiz.GetNodeId(ctx, param)
	if err != nil {
		return nil, err
	}

	cfg := s.nodeIDBiz.GetConfig()
	return &resourcev1.GetNodeIdResp{
		Data: dto.NodeIDDto.ToPbGetNodeIdRespData(cfg, dataModel),
	}, nil
}

// RenewalNodeId 续订节点id
func (s *nodeIDV1Service) RenewalNodeId(ctx context.Context, req *resourcev1.RenewalNodeIdReq) (*resourcev1.RenewalNodeIdResp, error) {
	//resp, err := s.renewalNodeIdWithQueue(ctx, req)
	//if err == nil {
	//	return resp, nil
	//}
	//s.log.WithContext(ctx).Errorw("msg", "RenewalNodeIdWithQueue failed!", "err", err)
	return s.renewalNodeId(ctx, req)
}

// RenewalNodeId 使用队列方式续订节点id
func (s *nodeIDV1Service) renewalNodeIdWithQueue(ctx context.Context, req *resourcev1.RenewalNodeIdReq) (*resourcev1.RenewalNodeIdResp, error) {
	var (
		param    = dto.NodeIDDto.ToBoRenewalNodeIdParam(req)
		conf     = s.nodeIDBiz.GetConfig()
		respData = &bo.RenewalNodeIDReply{
			Status:    enumv1.NodeIDStatusEnum_USING,
			ExpiredAt: conf.NextExpireTime(time.Now()),
		}
	)
	err := s.renewNodeIDEvent.Publish(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.RenewalNodeIdResp{
		Data: dto.NodeIDDto.ToPbRenewalNodeIdRespData2(respData),
	}, nil
}

// RenewalNodeId 续订节点id
func (s *nodeIDV1Service) renewalNodeId(ctx context.Context, req *resourcev1.RenewalNodeIdReq) (*resourcev1.RenewalNodeIdResp, error) {
	param := dto.NodeIDDto.ToBoRenewalNodeIdParam(req)
	dataModel, err := s.nodeIDBiz.RenewalNodeId(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.RenewalNodeIdResp{
		Data: dto.NodeIDDto.ToPbRenewalNodeIdRespData2(dataModel),
	}, nil
}

// ReleaseNodeId 释放节点id
func (s *nodeIDV1Service) ReleaseNodeId(ctx context.Context, req *resourcev1.ReleaseNodeIdReq) (*resourcev1.ReleaseNodeIdResp, error) {
	param := dto.NodeIDDto.ToBoReleaseNodeIdParam(req)
	dataModel, err := s.nodeIDBiz.ReleaseNodeId(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.ReleaseNodeIdResp{
		Data: dto.NodeIDDto.ToPbReleaseNodeIdRespData(dataModel),
	}, nil
}
