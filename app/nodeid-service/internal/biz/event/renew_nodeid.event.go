package events

import (
	"context"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/repo"
)

var (
	_ bizrepos.RenewNodeIDEventRepo = (*renewNodeIDEvent)(nil)
)

type renewNodeIDEvent struct {
	log           *log.Helper
	mqConn        *amqp.ConnectionWrapper
	nodeidBizRepo bizrepos.NodeIdBizRepo
}

func NewRenewNodeIDEventRepo(
	logger log.Logger,
	mqConn *amqp.ConnectionWrapper,
	nodeidBizRepo bizrepos.NodeIdBizRepo,
) bizrepos.RenewNodeIDEventRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "nodeid-service/biz/event/renew_nodeid"))
	return &renewNodeIDEvent{
		log:           logHelper,
		mqConn:        mqConn,
		nodeidBizRepo: nodeidBizRepo,
	}
}

func (s *renewNodeIDEvent) Send(ctx context.Context, param *bo.RenewalNodeIdParam) error {
	s.mqConn.IsConnected()
	return nil
}

func (s *renewNodeIDEvent) Receive(ctx context.Context, handler bizrepos.RenewNodeIDHandler) error {
	return nil
}

func (s *renewNodeIDEvent) Process(ctx context.Context, param *bo.RenewalNodeIdParam) error {
	_, err := s.nodeidBizRepo.RenewalNodeId(ctx, param)
	if err != nil {
		s.log.Infow("RenewalNodeId failed", "param", param, "err", err)
		return err
	}
	return nil
}

func (s *renewNodeIDEvent) Close(ctx context.Context) error {
	return nil
}
