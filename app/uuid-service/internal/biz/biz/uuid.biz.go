package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/uuid-service/internal/biz/repo"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type uuidBiz struct {
	log  *log.Helper
	node idpkg.Snowflake
}

func NewUuidBiz(
	logger log.Logger,
	node idpkg.Snowflake,
) bizrepos.UuidBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "uuid-service/biz/biz"))

	return &uuidBiz{
		log:  logHelper,
		node: node,
	}
}

func (s *uuidBiz) NextID(ctx context.Context) (uint64, error) {
	id, err := s.node.NextID()
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return 0, errorpkg.WithStack(e)
	}
	return id, nil
}
