package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	resourcev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	servicev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/services"
	errorutil "github.com/go-micro-saas/service-kit/error"
	clientpkg "github.com/ikaiguang/go-srv-kit/kratos/client"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

var (
	httpEndpoint = flag.String("http_endpoint", "localhost:10201", "HTTP endpoint")
	grpcEndpoint = flag.String("grpc_endpoint", "localhost:10202", "gRPC endpoint")
	instanceID   = flag.String("instance_id", "my_test_instance", "instance id")
	instanceName = flag.String("instance_name", "my_test_instance", "instance name")
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	handler := &getNodeID{}
	nodeIDClient, err := handler.GetNodeIDV1GRPCClient()
	if err != nil {
		fmt.Printf("==> GetNodeIDV1GRPCClient failed: %+v\n", err)
		return
	}
	serviceInfo, err := handler.GetServiceInfo(nodeIDClient)
	if err != nil {
		fmt.Printf("==> GetServiceInfo failed: %+v\n", err)
		return
	}
	fmt.Printf("==> GetServiceInfo result: %#v\n", serviceInfo.String())
	nodeIDData, err := handler.GetNodeId(nodeIDClient)
	if err != nil {
		fmt.Printf("==> GetNodeId failed: %+v\n", err)
		return
	}
	fmt.Printf("==> GetNodeId result: %#v\n", nodeIDData.String())
	renewalData, err := handler.RenewalNodeId(nodeIDClient, nodeIDData)
	if err != nil {
		fmt.Printf("==> RenewalNodeId failed: %+v\n", err)
		return
	}
	fmt.Printf("==> RenewalNodeId result: %#v\n", renewalData.String())
	releaseData, err := handler.ReleaseNodeId(nodeIDClient, nodeIDData)
	if err != nil {
		fmt.Printf("==> ReleaseNodeId failed: %+v\n", err)
		return
	}
	fmt.Printf("==> ReleaseNodeId result: %#v\n", releaseData.String())
}

type getNodeID struct{}

func (s *getNodeID) GetNodeIDV1GRPCClient() (servicev1.SrvNodeIDV1Client, error) {
	var opts = []grpc.ClientOption{
		grpc.WithEndpoint(*grpcEndpoint),
	}

	ctx := context.Background()

	clientCC, err := clientpkg.NewGRPCClient(ctx, true, opts...)
	if err != nil {
		return nil, errorpkg.FormatError(err)
	}
	return servicev1.NewSrvNodeIDV1Client(clientCC), nil
}

func (s *getNodeID) GetServiceInfo(grpcClient servicev1.SrvNodeIDV1Client) (*resourcev1.GetServiceInfoRespData, error) {
	getReq := &resourcev1.GetServiceInfoReq{}
	getResp, err := grpcClient.GetServiceInfo(context.Background(), getReq)
	if err != nil {
		return nil, errorpkg.FormatError(err)
	}
	if e := errorutil.CheckResponseCode(getResp); e != nil {
		return nil, errorpkg.FormatError(e)
	}
	return getResp.Data, nil
}

func (s *getNodeID) GetNodeId(grpcClient servicev1.SrvNodeIDV1Client) (*resourcev1.GetNodeIdRespData, error) {
	getReq := &resourcev1.GetNodeIdReq{
		InstanceId:   *instanceID,
		InstanceName: *instanceName,
		Metadata:     nil,
	}
	getResp, err := grpcClient.GetNodeId(context.Background(), getReq)
	if err != nil {
		return nil, errorpkg.FormatError(err)
	}
	if e := errorutil.CheckResponseCode(getResp); e != nil {
		return nil, errorpkg.FormatError(e)
	}
	return getResp.Data, nil
}

func (s *getNodeID) RenewalNodeId(grpcClient servicev1.SrvNodeIDV1Client, dataModel *resourcev1.GetNodeIdRespData) (*resourcev1.RenewalNodeIdRespData, error) {
	getReq := &resourcev1.RenewalNodeIdReq{
		Id:         dataModel.Id,
		InstanceId: dataModel.InstanceId,
		NodeId:     dataModel.NodeId,
	}
	getResp, err := grpcClient.RenewalNodeId(context.Background(), getReq)
	if err != nil {
		return nil, errorpkg.FormatError(err)
	}
	if e := errorutil.CheckResponseCode(getResp); e != nil {
		return nil, errorpkg.FormatError(e)
	}
	return getResp.Data, nil
}

func (s *getNodeID) ReleaseNodeId(grpcClient servicev1.SrvNodeIDV1Client, dataModel *resourcev1.GetNodeIdRespData) (*resourcev1.ReleaseNodeIdRespData, error) {
	getReq := &resourcev1.ReleaseNodeIdReq{
		Id:         dataModel.Id,
		InstanceId: dataModel.InstanceId,
		NodeId:     dataModel.NodeId,
	}
	getResp, err := grpcClient.ReleaseNodeId(context.Background(), getReq)
	if err != nil {
		return nil, errorpkg.FormatError(err)
	}
	if e := errorutil.CheckResponseCode(getResp); e != nil {
		return nil, errorpkg.FormatError(e)
	}
	return getResp.Data, nil
}
