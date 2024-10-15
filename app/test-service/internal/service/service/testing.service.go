package service

import (
	"github.com/go-kratos/kratos/v2/log"
	servicev1 "github.com/go-micro-saas/nodeid-service/api/test-service/v1/services"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/test-service/internal/biz/repo"
)

type testingV1Service struct {
	servicev1.UnimplementedSrvTestV1Server

	log        *log.Helper
	testingBiz bizrepos.TestingBizRepo
}

func NewTestingV1Service(logger log.Logger, testingBiz bizrepos.TestingBizRepo) servicev1.SrvTestV1Server {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/service/service"))
	return &testingV1Service{
		log:        logHelper,
		testingBiz: testingBiz,
	}
}
