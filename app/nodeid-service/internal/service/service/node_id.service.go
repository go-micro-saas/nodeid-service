package service

import (
	"context"
	resources "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	servicev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/services"
)

type nodeIDV1Service struct {
	servicev1.UnimplementedSrvNodeIDV1Server
}

// Ping ping
func (s *nodeIDV1Service) Ping(context.Context, *resources.PingReq) (*resources.PingResp, error) {
	return &resources.PingResp{}, nil
}

// GetServiceInfo 获取服务信息
func (s *nodeIDV1Service) GetServiceInfo(context.Context, *resources.GetServiceInfoReq) (*resources.GetServiceInfoResp, error) {
	return &resources.GetServiceInfoResp{}, nil
}

// GetNodeId 获取节点id
func (s *nodeIDV1Service) GetNodeId(context.Context, *resources.GetNodeIdReq) (*resources.GetNodeIdResp, error) {
	return &resources.GetNodeIdResp{}, nil
}

// RenewalNodeId 续订节点id
func (s *nodeIDV1Service) RenewalNodeId(context.Context, *resources.RenewalNodeIdReq) (*resources.RenewalNodeIdResp, error) {
	return &resources.RenewalNodeIdResp{}, nil
}

// ReleaseNodeId 释放节点id
func (s *nodeIDV1Service) ReleaseNodeId(context.Context, *resources.ReleaseNodeIdReq) (*resources.ReleaseNodeIdResp, error) {
	return &resources.ReleaseNodeIdResp{}, nil
}
