package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/repo"
	datarepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/repo"
)

type nodeIDBiz struct {
	log        *log.Helper
	conf       *bo.NodeIDConfig
	nodeIDData datarepos.NodeIdDataRepo
}

func NewNodeIDBiz(
	logger log.Logger,
	conf *bo.NodeIDConfig,
	nodeIDData datarepos.NodeIdDataRepo,
) bizrepos.NodeIdBizRepo {
	conf.Initialization()

	logHelper := log.NewHelper(log.With(logger, "module", "nodeid-service/biz/biz"))
	return &nodeIDBiz{
		log:        logHelper,
		conf:       conf,
		nodeIDData: nodeIDData,
	}
}

func (s *nodeIDBiz) GetConfig() *bo.NodeIDConfig {
	cfg := *s.conf
	return &cfg
}
