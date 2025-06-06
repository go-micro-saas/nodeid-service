package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	resourcev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	servicev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/services"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/repo"
	threadpkg "github.com/ikaiguang/go-srv-kit/kratos/thread"
	"sync/atomic"
)

type nodeEventV1Service struct {
	servicev1.UnimplementedSrvNodeEventV1Server

	log              *log.Helper
	nodeIDBiz        bizrepos.NodeIdBizRepo
	renewNodeIDEvent bizrepos.RenewNodeIDEventRepo

	isConsumingRenewNodeIDEvent atomic.Int64
}

func NewNodeEventV1Service(
	logger log.Logger,
	nodeIDBiz bizrepos.NodeIdBizRepo,
	renewNodeIDEvent bizrepos.RenewNodeIDEventRepo,
) servicev1.SrvNodeEventV1Server {
	logHelper := log.NewHelper(log.With(logger, "module", "nodeid-service/service/service"))
	return &nodeEventV1Service{
		log:       logHelper,
		nodeIDBiz: nodeIDBiz,
		//renewNodeIDEvent: renewNodeIDEvent,
	}
}

// Ping ping
func (s *nodeEventV1Service) Ping(ctx context.Context, req *resourcev1.PingReq) (*resourcev1.PingResp, error) {
	return &resourcev1.PingResp{Data: &resourcev1.PingRespData{
		Message: "request message: " + req.Message,
	}}, nil
}

// SubscribeRenewalNodeIdEvent 订阅续订节点id事件
func (s *nodeEventV1Service) SubscribeRenewalNodeIdEvent(ctx context.Context, req *resourcev1.SubscribeRenewalNodeIdEventReq) (*resourcev1.SubscribeRenewalNodeIdEventResp, error) {
	if s.isConsumingRenewNodeIDEvent.Load() < 1 {
		threadpkg.GoSafe(func() {
			s.isConsumingRenewNodeIDEvent.Add(1)
			defer func() { s.isConsumingRenewNodeIDEvent.Add(-1) }()
			err := s.renewNodeIDEvent.Consume(ctx, s.nodeIDBiz.RenewalNodeId)
			if err != nil {
				s.log.WithContext(ctx).Errorw("msg", "SubscribeRenewalNodeIdEvent failed!", "err", err)
			}
		})
	}
	return &resourcev1.SubscribeRenewalNodeIdEventResp{
		Data: &resourcev1.SubscribeRenewalNodeIdEventRespData{
			ConsumerCounter: s.isConsumingRenewNodeIDEvent.Load(),
		},
	}, nil
}

// StopRenewalNodeIdEvent 停止续订节点id事件
func (s *nodeEventV1Service) StopRenewalNodeIdEvent(ctx context.Context, req *resourcev1.StopRenewalNodeIdEventReq) (*resourcev1.StopRenewalNodeIdEventResp, error) {
	err := s.renewNodeIDEvent.Close(ctx)
	if err != nil {
		s.log.WithContext(ctx).Errorw("msg", "run StopRenewalNodeIdEvent failed!", "err", err)
		return nil, err
	}
	return &resourcev1.StopRenewalNodeIdEventResp{
		Data: &resourcev1.StopRenewalNodeIdEventRespData{
			ConsumerCounter: s.isConsumingRenewNodeIDEvent.Load(),
		},
	}, nil
}
