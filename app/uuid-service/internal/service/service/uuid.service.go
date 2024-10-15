package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	resourcev1 "github.com/go-micro-saas/nodeid-service/api/uuid-service/v1/resources"
	servicev1 "github.com/go-micro-saas/nodeid-service/api/uuid-service/v1/services"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/biz/repo"
)

type uuidV1Service struct {
	servicev1.UnimplementedSrvUuidV1Server

	log     *log.Helper
	uuidBiz bizrepos.UuidBizRepo
}

func NewUuidV1Service(logger log.Logger, uuidBiz bizrepos.UuidBizRepo) servicev1.SrvUuidV1Server {
	logHelper := log.NewHelper(log.With(logger, "module", "uuid-service/service/service"))
	return &uuidV1Service{
		log:     logHelper,
		uuidBiz: uuidBiz,
	}
}

func (s *uuidV1Service) NextID(ctx context.Context, req *resourcev1.NextIDReq) (*resourcev1.NextIDResp, error) {
	return s.UnimplementedSrvUuidV1Server.NextID(ctx, req)
}
