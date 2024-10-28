package dto

import (
	nodeidresourcev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	resourcev1 "github.com/go-micro-saas/nodeid-service/api/uuid-service/v1/resources"
	"github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/conf"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	UuidDto uuidDto
)

type uuidDto struct{}

func ToPbGetNodeIdReq(cfg *conf.ServiceConfig) *nodeidresourcev1.GetNodeIdReq {
	res := &nodeidresourcev1.GetNodeIdReq{}
	if cfg.GetUuidService() == nil {
		return res
	}
	serviceConfig := cfg.GetUuidService()
	res.InstanceId = serviceConfig.InstanceId
	res.InstanceName = serviceConfig.InstanceName
	res.Metadata = serviceConfig.Metadata
	return res
}

func (s *uuidDto) ToBoXxx() interface{} {
	return emptypb.Empty{}
}

func (s *uuidDto) ToPbNextIDRespData(id uint64) *resourcev1.NextIDRespData {
	res := &resourcev1.NextIDRespData{
		Id: id,
	}
	return res
}
