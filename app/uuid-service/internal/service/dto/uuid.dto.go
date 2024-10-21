package dto

import (
	resourcev1 "github.com/go-micro-saas/nodeid-service/api/uuid-service/v1/resources"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	UuidDto uuidDto
)

type uuidDto struct{}

func (s *uuidDto) ToBoXxx() interface{} {
	return emptypb.Empty{}
}

func (s *uuidDto) ToPbNextIDRespData(id uint64) *resourcev1.NextIDRespData {
	res := &resourcev1.NextIDRespData{
		Id: id,
	}
	return res
}
