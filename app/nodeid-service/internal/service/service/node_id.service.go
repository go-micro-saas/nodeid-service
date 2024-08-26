package service

import (
	"context"
	resourcev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	servicev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/services"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
)

type nodeIDV1Service struct {
	servicev1.UnimplementedSrvNodeIDV1Server

	conf *bo.NodeIDConfig
}

func NewNodeIDV1Service(conf *bo.NodeIDConfig) servicev1.SrvNodeIDV1Server {
	conf.Initialization()
	return &nodeIDV1Service{
		conf: conf,
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
	return &resourcev1.GetServiceInfoResp{}, nil
}

// GetNodeId 获取节点id
func (s *nodeIDV1Service) GetNodeId(ctx context.Context, req *resourcev1.GetNodeIdReq) (*resourcev1.GetNodeIdResp, error) {
	return &resourcev1.GetNodeIdResp{}, nil
}

// RenewalNodeId 续订节点id
func (s *nodeIDV1Service) RenewalNodeId(ctx context.Context, req *resourcev1.RenewalNodeIdReq) (*resourcev1.RenewalNodeIdResp, error) {
	return &resourcev1.RenewalNodeIdResp{}, nil
}

// ReleaseNodeId 释放节点id
func (s *nodeIDV1Service) ReleaseNodeId(ctx context.Context, req *resourcev1.ReleaseNodeIdReq) (*resourcev1.ReleaseNodeIdResp, error) {
	return &resourcev1.ReleaseNodeIdResp{}, nil
}
