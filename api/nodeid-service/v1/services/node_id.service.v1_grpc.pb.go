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
	SrvTestV1_Ping_FullMethodName           = "/saas.api.nodeid.servicev1.SrvTestV1/Ping"
	SrvTestV1_GetServiceInfo_FullMethodName = "/saas.api.nodeid.servicev1.SrvTestV1/GetServiceInfo"
	SrvTestV1_GetNodeId_FullMethodName      = "/saas.api.nodeid.servicev1.SrvTestV1/GetNodeId"
	SrvTestV1_RenewalNodeId_FullMethodName  = "/saas.api.nodeid.servicev1.SrvTestV1/RenewalNodeId"
	SrvTestV1_ReleaseNodeId_FullMethodName  = "/saas.api.nodeid.servicev1.SrvTestV1/ReleaseNodeId"
)

// SrvTestV1Client is the client API for SrvTestV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvTestV1Client interface {
	// ping
	Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error)
	// 获取服务信息
	GetServiceInfo(ctx context.Context, in *resources.GetServiceInfoReq, opts ...grpc.CallOption) (*resources.GetServiceInfoResp, error)
	// 获取节点id
	GetNodeId(ctx context.Context, in *resources.GetNodeIdReq, opts ...grpc.CallOption) (*resources.GetNodeIdResp, error)
	// 续订节点id
	RenewalNodeId(ctx context.Context, in *resources.RenewalNodeIdReq, opts ...grpc.CallOption) (*resources.RenewalNodeIdResp, error)
	// 释放节点id
	ReleaseNodeId(ctx context.Context, in *resources.ReleaseNodeIdReq, opts ...grpc.CallOption) (*resources.ReleaseNodeIdResp, error)
}

type srvTestV1Client struct {
	cc grpc.ClientConnInterface
}

func NewSrvTestV1Client(cc grpc.ClientConnInterface) SrvTestV1Client {
	return &srvTestV1Client{cc}
}

func (c *srvTestV1Client) Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error) {
	out := new(resources.PingResp)
	err := c.cc.Invoke(ctx, SrvTestV1_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvTestV1Client) GetServiceInfo(ctx context.Context, in *resources.GetServiceInfoReq, opts ...grpc.CallOption) (*resources.GetServiceInfoResp, error) {
	out := new(resources.GetServiceInfoResp)
	err := c.cc.Invoke(ctx, SrvTestV1_GetServiceInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvTestV1Client) GetNodeId(ctx context.Context, in *resources.GetNodeIdReq, opts ...grpc.CallOption) (*resources.GetNodeIdResp, error) {
	out := new(resources.GetNodeIdResp)
	err := c.cc.Invoke(ctx, SrvTestV1_GetNodeId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvTestV1Client) RenewalNodeId(ctx context.Context, in *resources.RenewalNodeIdReq, opts ...grpc.CallOption) (*resources.RenewalNodeIdResp, error) {
	out := new(resources.RenewalNodeIdResp)
	err := c.cc.Invoke(ctx, SrvTestV1_RenewalNodeId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvTestV1Client) ReleaseNodeId(ctx context.Context, in *resources.ReleaseNodeIdReq, opts ...grpc.CallOption) (*resources.ReleaseNodeIdResp, error) {
	out := new(resources.ReleaseNodeIdResp)
	err := c.cc.Invoke(ctx, SrvTestV1_ReleaseNodeId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvTestV1Server is the server API for SrvTestV1 service.
// All implementations must embed UnimplementedSrvTestV1Server
// for forward compatibility
type SrvTestV1Server interface {
	// ping
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
	// 获取服务信息
	GetServiceInfo(context.Context, *resources.GetServiceInfoReq) (*resources.GetServiceInfoResp, error)
	// 获取节点id
	GetNodeId(context.Context, *resources.GetNodeIdReq) (*resources.GetNodeIdResp, error)
	// 续订节点id
	RenewalNodeId(context.Context, *resources.RenewalNodeIdReq) (*resources.RenewalNodeIdResp, error)
	// 释放节点id
	ReleaseNodeId(context.Context, *resources.ReleaseNodeIdReq) (*resources.ReleaseNodeIdResp, error)
	mustEmbedUnimplementedSrvTestV1Server()
}

// UnimplementedSrvTestV1Server must be embedded to have forward compatible implementations.
type UnimplementedSrvTestV1Server struct {
}

func (UnimplementedSrvTestV1Server) Ping(context.Context, *resources.PingReq) (*resources.PingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSrvTestV1Server) GetServiceInfo(context.Context, *resources.GetServiceInfoReq) (*resources.GetServiceInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceInfo not implemented")
}
func (UnimplementedSrvTestV1Server) GetNodeId(context.Context, *resources.GetNodeIdReq) (*resources.GetNodeIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeId not implemented")
}
func (UnimplementedSrvTestV1Server) RenewalNodeId(context.Context, *resources.RenewalNodeIdReq) (*resources.RenewalNodeIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewalNodeId not implemented")
}
func (UnimplementedSrvTestV1Server) ReleaseNodeId(context.Context, *resources.ReleaseNodeIdReq) (*resources.ReleaseNodeIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseNodeId not implemented")
}
func (UnimplementedSrvTestV1Server) mustEmbedUnimplementedSrvTestV1Server() {}

// UnsafeSrvTestV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvTestV1Server will
// result in compilation errors.
type UnsafeSrvTestV1Server interface {
	mustEmbedUnimplementedSrvTestV1Server()
}

func RegisterSrvTestV1Server(s grpc.ServiceRegistrar, srv SrvTestV1Server) {
	s.RegisterService(&SrvTestV1_ServiceDesc, srv)
}

func _SrvTestV1_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvTestV1Server).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvTestV1_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvTestV1Server).Ping(ctx, req.(*resources.PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvTestV1_GetServiceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetServiceInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvTestV1Server).GetServiceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvTestV1_GetServiceInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvTestV1Server).GetServiceInfo(ctx, req.(*resources.GetServiceInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvTestV1_GetNodeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetNodeIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvTestV1Server).GetNodeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvTestV1_GetNodeId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvTestV1Server).GetNodeId(ctx, req.(*resources.GetNodeIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvTestV1_RenewalNodeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.RenewalNodeIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvTestV1Server).RenewalNodeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvTestV1_RenewalNodeId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvTestV1Server).RenewalNodeId(ctx, req.(*resources.RenewalNodeIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvTestV1_ReleaseNodeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.ReleaseNodeIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvTestV1Server).ReleaseNodeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvTestV1_ReleaseNodeId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvTestV1Server).ReleaseNodeId(ctx, req.(*resources.ReleaseNodeIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvTestV1_ServiceDesc is the grpc.ServiceDesc for SrvTestV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvTestV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "saas.api.nodeid.servicev1.SrvTestV1",
	HandlerType: (*SrvTestV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SrvTestV1_Ping_Handler,
		},
		{
			MethodName: "GetServiceInfo",
			Handler:    _SrvTestV1_GetServiceInfo_Handler,
		},
		{
			MethodName: "GetNodeId",
			Handler:    _SrvTestV1_GetNodeId_Handler,
		},
		{
			MethodName: "RenewalNodeId",
			Handler:    _SrvTestV1_RenewalNodeId_Handler,
		},
		{
			MethodName: "ReleaseNodeId",
			Handler:    _SrvTestV1_ReleaseNodeId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/nodeid-service/v1/services/node_id.service.v1.proto",
}