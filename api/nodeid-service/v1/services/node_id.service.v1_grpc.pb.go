// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: api/nodeid-service/v1/services/node_id.service.v1.proto

package servicev1

import (
	context "context"
	resources "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SrvNodeIDV1_Ping_FullMethodName                        = "/saas.api.nodeid.servicev1.SrvNodeIDV1/Ping"
	SrvNodeIDV1_GetServiceInfo_FullMethodName              = "/saas.api.nodeid.servicev1.SrvNodeIDV1/GetServiceInfo"
	SrvNodeIDV1_GetNodeId_FullMethodName                   = "/saas.api.nodeid.servicev1.SrvNodeIDV1/GetNodeId"
	SrvNodeIDV1_RenewalNodeId_FullMethodName               = "/saas.api.nodeid.servicev1.SrvNodeIDV1/RenewalNodeId"
	SrvNodeIDV1_ReleaseNodeId_FullMethodName               = "/saas.api.nodeid.servicev1.SrvNodeIDV1/ReleaseNodeId"
	SrvNodeIDV1_SubscribeRenewalNodeIdEvent_FullMethodName = "/saas.api.nodeid.servicev1.SrvNodeIDV1/SubscribeRenewalNodeIdEvent"
	SrvNodeIDV1_StopRenewalNodeIdEvent_FullMethodName      = "/saas.api.nodeid.servicev1.SrvNodeIDV1/StopRenewalNodeIdEvent"
)

// SrvNodeIDV1Client is the client API for SrvNodeIDV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvNodeIDV1Client interface {
	// ping
	Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error)
	// 节点ID - 获取服务信息
	GetServiceInfo(ctx context.Context, in *resources.GetServiceInfoReq, opts ...grpc.CallOption) (*resources.GetServiceInfoResp, error)
	// 节点ID - 获取节点id
	GetNodeId(ctx context.Context, in *resources.GetNodeIdReq, opts ...grpc.CallOption) (*resources.GetNodeIdResp, error)
	// 节点ID - 续订节点id
	RenewalNodeId(ctx context.Context, in *resources.RenewalNodeIdReq, opts ...grpc.CallOption) (*resources.RenewalNodeIdResp, error)
	// 节点ID - 释放节点id
	ReleaseNodeId(ctx context.Context, in *resources.ReleaseNodeIdReq, opts ...grpc.CallOption) (*resources.ReleaseNodeIdResp, error)
	// 节点ID - 订阅续订节点id事件
	SubscribeRenewalNodeIdEvent(ctx context.Context, in *resources.SubscribeRenewalNodeIdEventReq, opts ...grpc.CallOption) (*resources.SubscribeRenewalNodeIdEventResp, error)
	// 节点ID - 停止续订节点id事件
	StopRenewalNodeIdEvent(ctx context.Context, in *resources.StopRenewalNodeIdEventReq, opts ...grpc.CallOption) (*resources.StopRenewalNodeIdEventResp, error)
}

type srvNodeIDV1Client struct {
	cc grpc.ClientConnInterface
}

func NewSrvNodeIDV1Client(cc grpc.ClientConnInterface) SrvNodeIDV1Client {
	return &srvNodeIDV1Client{cc}
}

func (c *srvNodeIDV1Client) Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error) {
	out := new(resources.PingResp)
	err := c.cc.Invoke(ctx, SrvNodeIDV1_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvNodeIDV1Client) GetServiceInfo(ctx context.Context, in *resources.GetServiceInfoReq, opts ...grpc.CallOption) (*resources.GetServiceInfoResp, error) {
	out := new(resources.GetServiceInfoResp)
	err := c.cc.Invoke(ctx, SrvNodeIDV1_GetServiceInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvNodeIDV1Client) GetNodeId(ctx context.Context, in *resources.GetNodeIdReq, opts ...grpc.CallOption) (*resources.GetNodeIdResp, error) {
	out := new(resources.GetNodeIdResp)
	err := c.cc.Invoke(ctx, SrvNodeIDV1_GetNodeId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvNodeIDV1Client) RenewalNodeId(ctx context.Context, in *resources.RenewalNodeIdReq, opts ...grpc.CallOption) (*resources.RenewalNodeIdResp, error) {
	out := new(resources.RenewalNodeIdResp)
	err := c.cc.Invoke(ctx, SrvNodeIDV1_RenewalNodeId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvNodeIDV1Client) ReleaseNodeId(ctx context.Context, in *resources.ReleaseNodeIdReq, opts ...grpc.CallOption) (*resources.ReleaseNodeIdResp, error) {
	out := new(resources.ReleaseNodeIdResp)
	err := c.cc.Invoke(ctx, SrvNodeIDV1_ReleaseNodeId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvNodeIDV1Client) SubscribeRenewalNodeIdEvent(ctx context.Context, in *resources.SubscribeRenewalNodeIdEventReq, opts ...grpc.CallOption) (*resources.SubscribeRenewalNodeIdEventResp, error) {
	out := new(resources.SubscribeRenewalNodeIdEventResp)
	err := c.cc.Invoke(ctx, SrvNodeIDV1_SubscribeRenewalNodeIdEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvNodeIDV1Client) StopRenewalNodeIdEvent(ctx context.Context, in *resources.StopRenewalNodeIdEventReq, opts ...grpc.CallOption) (*resources.StopRenewalNodeIdEventResp, error) {
	out := new(resources.StopRenewalNodeIdEventResp)
	err := c.cc.Invoke(ctx, SrvNodeIDV1_StopRenewalNodeIdEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvNodeIDV1Server is the server API for SrvNodeIDV1 service.
// All implementations must embed UnimplementedSrvNodeIDV1Server
// for forward compatibility
type SrvNodeIDV1Server interface {
	// ping
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
	// 节点ID - 获取服务信息
	GetServiceInfo(context.Context, *resources.GetServiceInfoReq) (*resources.GetServiceInfoResp, error)
	// 节点ID - 获取节点id
	GetNodeId(context.Context, *resources.GetNodeIdReq) (*resources.GetNodeIdResp, error)
	// 节点ID - 续订节点id
	RenewalNodeId(context.Context, *resources.RenewalNodeIdReq) (*resources.RenewalNodeIdResp, error)
	// 节点ID - 释放节点id
	ReleaseNodeId(context.Context, *resources.ReleaseNodeIdReq) (*resources.ReleaseNodeIdResp, error)
	// 节点ID - 订阅续订节点id事件
	SubscribeRenewalNodeIdEvent(context.Context, *resources.SubscribeRenewalNodeIdEventReq) (*resources.SubscribeRenewalNodeIdEventResp, error)
	// 节点ID - 停止续订节点id事件
	StopRenewalNodeIdEvent(context.Context, *resources.StopRenewalNodeIdEventReq) (*resources.StopRenewalNodeIdEventResp, error)
	mustEmbedUnimplementedSrvNodeIDV1Server()
}

// UnimplementedSrvNodeIDV1Server must be embedded to have forward compatible implementations.
type UnimplementedSrvNodeIDV1Server struct {
}

func (UnimplementedSrvNodeIDV1Server) Ping(context.Context, *resources.PingReq) (*resources.PingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSrvNodeIDV1Server) GetServiceInfo(context.Context, *resources.GetServiceInfoReq) (*resources.GetServiceInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceInfo not implemented")
}
func (UnimplementedSrvNodeIDV1Server) GetNodeId(context.Context, *resources.GetNodeIdReq) (*resources.GetNodeIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeId not implemented")
}
func (UnimplementedSrvNodeIDV1Server) RenewalNodeId(context.Context, *resources.RenewalNodeIdReq) (*resources.RenewalNodeIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewalNodeId not implemented")
}
func (UnimplementedSrvNodeIDV1Server) ReleaseNodeId(context.Context, *resources.ReleaseNodeIdReq) (*resources.ReleaseNodeIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseNodeId not implemented")
}
func (UnimplementedSrvNodeIDV1Server) SubscribeRenewalNodeIdEvent(context.Context, *resources.SubscribeRenewalNodeIdEventReq) (*resources.SubscribeRenewalNodeIdEventResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeRenewalNodeIdEvent not implemented")
}
func (UnimplementedSrvNodeIDV1Server) StopRenewalNodeIdEvent(context.Context, *resources.StopRenewalNodeIdEventReq) (*resources.StopRenewalNodeIdEventResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopRenewalNodeIdEvent not implemented")
}
func (UnimplementedSrvNodeIDV1Server) mustEmbedUnimplementedSrvNodeIDV1Server() {}

// UnsafeSrvNodeIDV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvNodeIDV1Server will
// result in compilation errors.
type UnsafeSrvNodeIDV1Server interface {
	mustEmbedUnimplementedSrvNodeIDV1Server()
}

func RegisterSrvNodeIDV1Server(s grpc.ServiceRegistrar, srv SrvNodeIDV1Server) {
	s.RegisterService(&SrvNodeIDV1_ServiceDesc, srv)
}

func _SrvNodeIDV1_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvNodeIDV1Server).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvNodeIDV1_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvNodeIDV1Server).Ping(ctx, req.(*resources.PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvNodeIDV1_GetServiceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetServiceInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvNodeIDV1Server).GetServiceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvNodeIDV1_GetServiceInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvNodeIDV1Server).GetServiceInfo(ctx, req.(*resources.GetServiceInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvNodeIDV1_GetNodeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetNodeIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvNodeIDV1Server).GetNodeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvNodeIDV1_GetNodeId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvNodeIDV1Server).GetNodeId(ctx, req.(*resources.GetNodeIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvNodeIDV1_RenewalNodeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.RenewalNodeIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvNodeIDV1Server).RenewalNodeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvNodeIDV1_RenewalNodeId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvNodeIDV1Server).RenewalNodeId(ctx, req.(*resources.RenewalNodeIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvNodeIDV1_ReleaseNodeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.ReleaseNodeIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvNodeIDV1Server).ReleaseNodeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvNodeIDV1_ReleaseNodeId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvNodeIDV1Server).ReleaseNodeId(ctx, req.(*resources.ReleaseNodeIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvNodeIDV1_SubscribeRenewalNodeIdEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.SubscribeRenewalNodeIdEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvNodeIDV1Server).SubscribeRenewalNodeIdEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvNodeIDV1_SubscribeRenewalNodeIdEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvNodeIDV1Server).SubscribeRenewalNodeIdEvent(ctx, req.(*resources.SubscribeRenewalNodeIdEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvNodeIDV1_StopRenewalNodeIdEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.StopRenewalNodeIdEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvNodeIDV1Server).StopRenewalNodeIdEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvNodeIDV1_StopRenewalNodeIdEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvNodeIDV1Server).StopRenewalNodeIdEvent(ctx, req.(*resources.StopRenewalNodeIdEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvNodeIDV1_ServiceDesc is the grpc.ServiceDesc for SrvNodeIDV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvNodeIDV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "saas.api.nodeid.servicev1.SrvNodeIDV1",
	HandlerType: (*SrvNodeIDV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SrvNodeIDV1_Ping_Handler,
		},
		{
			MethodName: "GetServiceInfo",
			Handler:    _SrvNodeIDV1_GetServiceInfo_Handler,
		},
		{
			MethodName: "GetNodeId",
			Handler:    _SrvNodeIDV1_GetNodeId_Handler,
		},
		{
			MethodName: "RenewalNodeId",
			Handler:    _SrvNodeIDV1_RenewalNodeId_Handler,
		},
		{
			MethodName: "ReleaseNodeId",
			Handler:    _SrvNodeIDV1_ReleaseNodeId_Handler,
		},
		{
			MethodName: "SubscribeRenewalNodeIdEvent",
			Handler:    _SrvNodeIDV1_SubscribeRenewalNodeIdEvent_Handler,
		},
		{
			MethodName: "StopRenewalNodeIdEvent",
			Handler:    _SrvNodeIDV1_StopRenewalNodeIdEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/nodeid-service/v1/services/node_id.service.v1.proto",
}
