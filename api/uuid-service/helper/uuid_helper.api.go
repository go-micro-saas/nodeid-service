package uuidapi

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	nodeidapi "github.com/go-micro-saas/nodeid-service/api/nodeid-service/helper"
	nodeidresourcev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type uuidHelper struct {
	log          *log.Helper
	nodeidHekper nodeidapi.NodeIDHelper
}

func NewUuidHelper(
	logger log.Logger,
	nodeidHekper nodeidapi.NodeIDHelper,
) UuidHelper {
	logHelper := log.NewHelper(log.With(logger, "module", "uuid-api"))
	return &uuidHelper{
		log:          logHelper,
		nodeidHekper: nodeidHekper,
	}
}

func (s *uuidHelper) GetSnowflakeNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error) {
	nodeID, renewal, err := s.nodeidHekper.GetAndAutoRenewalNodeID(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	//idpkg.DefaultEpoch = ""
	node, err := idpkg.NewBwmarrinSnowflake(nodeID.NodeId)
	if err != nil {
		_ = renewal.Stop(ctx)
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, nil, errorpkg.WithStack(e)
	}
	cleanup := func() {
		cleanContext := context.Background()
		cleanupErr := renewal.Stop(cleanContext)
		if cleanupErr != nil {
			s.log.WithContext(ctx).Warnw("msg", "[release] stop renewal failed", "err", cleanupErr)
		}
		_, cleanupErr = s.nodeidHekper.ReleaseNodeId(cleanContext, nodeID)
		if cleanupErr != nil {
			s.log.WithContext(ctx).Warnw("msg", "[release] release node failed", "err", cleanupErr)
		}
	}
	return node, cleanup, nil
}
