// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.21.6
// source: api/nodeid-service/v1/services/node_id.service.v1.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	resources "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvTestV1GetNodeId = "/saas.api.nodeid.servicev1.SrvTestV1/GetNodeId"
const OperationSrvTestV1GetServiceInfo = "/saas.api.nodeid.servicev1.SrvTestV1/GetServiceInfo"
const OperationSrvTestV1Ping = "/saas.api.nodeid.servicev1.SrvTestV1/Ping"
const OperationSrvTestV1ReleaseNodeId = "/saas.api.nodeid.servicev1.SrvTestV1/ReleaseNodeId"
const OperationSrvTestV1RenewalNodeId = "/saas.api.nodeid.servicev1.SrvTestV1/RenewalNodeId"

type SrvTestV1HTTPServer interface {
	// GetNodeId 获取节点id
	GetNodeId(context.Context, *resources.GetNodeIdReq) (*resources.GetNodeIdResp, error)
	// GetServiceInfo 获取服务信息
	GetServiceInfo(context.Context, *resources.GetServiceInfoReq) (*resources.GetServiceInfoResp, error)
	// Ping ping
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
	// ReleaseNodeId 释放节点id
	ReleaseNodeId(context.Context, *resources.ReleaseNodeIdReq) (*resources.ReleaseNodeIdResp, error)
	// RenewalNodeId 续订节点id
	RenewalNodeId(context.Context, *resources.RenewalNodeIdReq) (*resources.RenewalNodeIdResp, error)
}

func RegisterSrvTestV1HTTPServer(s *http.Server, srv SrvTestV1HTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/nodeid/ping/{message}", _SrvTestV1_Ping1_HTTP_Handler(srv))
	r.GET("/api/v1/nodeid/get-service-info", _SrvTestV1_GetServiceInfo0_HTTP_Handler(srv))
	r.GET("/api/v1/nodeid/get-one-node-id", _SrvTestV1_GetNodeId0_HTTP_Handler(srv))
	r.POST("/api/v1/nodeid/renewal-node-id", _SrvTestV1_RenewalNodeId0_HTTP_Handler(srv))
	r.PUT("/api/v1/nodeid/release-node-id", _SrvTestV1_ReleaseNodeId0_HTTP_Handler(srv))
}

func _SrvTestV1_Ping1_HTTP_Handler(srv SrvTestV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.PingReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvTestV1Ping)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*resources.PingReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.PingResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestV1_GetServiceInfo0_HTTP_Handler(srv SrvTestV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.GetServiceInfoReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvTestV1GetServiceInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetServiceInfo(ctx, req.(*resources.GetServiceInfoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.GetServiceInfoResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestV1_GetNodeId0_HTTP_Handler(srv SrvTestV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.GetNodeIdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvTestV1GetNodeId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNodeId(ctx, req.(*resources.GetNodeIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.GetNodeIdResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestV1_RenewalNodeId0_HTTP_Handler(srv SrvTestV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.RenewalNodeIdReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvTestV1RenewalNodeId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RenewalNodeId(ctx, req.(*resources.RenewalNodeIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.RenewalNodeIdResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestV1_ReleaseNodeId0_HTTP_Handler(srv SrvTestV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.ReleaseNodeIdReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvTestV1ReleaseNodeId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReleaseNodeId(ctx, req.(*resources.ReleaseNodeIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.ReleaseNodeIdResp)
		return ctx.Result(200, reply)
	}
}

type SrvTestV1HTTPClient interface {
	GetNodeId(ctx context.Context, req *resources.GetNodeIdReq, opts ...http.CallOption) (rsp *resources.GetNodeIdResp, err error)
	GetServiceInfo(ctx context.Context, req *resources.GetServiceInfoReq, opts ...http.CallOption) (rsp *resources.GetServiceInfoResp, err error)
	Ping(ctx context.Context, req *resources.PingReq, opts ...http.CallOption) (rsp *resources.PingResp, err error)
	ReleaseNodeId(ctx context.Context, req *resources.ReleaseNodeIdReq, opts ...http.CallOption) (rsp *resources.ReleaseNodeIdResp, err error)
	RenewalNodeId(ctx context.Context, req *resources.RenewalNodeIdReq, opts ...http.CallOption) (rsp *resources.RenewalNodeIdResp, err error)
}

type SrvTestV1HTTPClientImpl struct {
	cc *http.Client
}

func NewSrvTestV1HTTPClient(client *http.Client) SrvTestV1HTTPClient {
	return &SrvTestV1HTTPClientImpl{client}
}

func (c *SrvTestV1HTTPClientImpl) GetNodeId(ctx context.Context, in *resources.GetNodeIdReq, opts ...http.CallOption) (*resources.GetNodeIdResp, error) {
	var out resources.GetNodeIdResp
	pattern := "/api/v1/nodeid/get-one-node-id"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvTestV1GetNodeId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SrvTestV1HTTPClientImpl) GetServiceInfo(ctx context.Context, in *resources.GetServiceInfoReq, opts ...http.CallOption) (*resources.GetServiceInfoResp, error) {
	var out resources.GetServiceInfoResp
	pattern := "/api/v1/nodeid/get-service-info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvTestV1GetServiceInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SrvTestV1HTTPClientImpl) Ping(ctx context.Context, in *resources.PingReq, opts ...http.CallOption) (*resources.PingResp, error) {
	var out resources.PingResp
	pattern := "/api/v1/nodeid/ping/{message}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvTestV1Ping))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SrvTestV1HTTPClientImpl) ReleaseNodeId(ctx context.Context, in *resources.ReleaseNodeIdReq, opts ...http.CallOption) (*resources.ReleaseNodeIdResp, error) {
	var out resources.ReleaseNodeIdResp
	pattern := "/api/v1/nodeid/release-node-id"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvTestV1ReleaseNodeId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SrvTestV1HTTPClientImpl) RenewalNodeId(ctx context.Context, in *resources.RenewalNodeIdReq, opts ...http.CallOption) (*resources.RenewalNodeIdResp, error) {
	var out resources.RenewalNodeIdResp
	pattern := "/api/v1/nodeid/renewal-node-id"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvTestV1RenewalNodeId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
