package nodeidapi

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	nodeidresourcev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	"time"
)

type RenewalManager interface {
	Stop(ctx context.Context) error
	RenewalResult(ctx context.Context) *RenewalResult
}

type RenewalResult struct {
	Data     *nodeidresourcev1.RenewalNodeIdRespData
	Err      error
	LastTime time.Time
}

type renewalManager struct {
	data *RenewalResult
	stop func()
}

func (s *renewalManager) Stop(_ context.Context) error {
	if s.stop != nil {
		s.stop()
	}
	return nil
}

func (s *renewalManager) RenewalResult(_ context.Context) *RenewalResult {
	return s.data
}

type NodeIDHelper interface {
	GetAndAutoRenewalNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, RenewalManager, error)
	GetNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error)
	RenewalNodeID(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (RenewalManager, error)
	ReleaseNodeId(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (*nodeidresourcev1.ReleaseNodeIdRespData, error)
}

type Option func(*options)

func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

func WithTries(tries int) Option {
	return func(o *options) {
		o.tries = tries
	}
}

func WithRetryDelay(delay time.Duration) Option {
	return func(o *options) {
		o.retryDelay = delay
	}
}

func WithHeartbeatInterval(duration time.Duration) Option {
	return func(o *options) {
		o.heartbeatInterval = duration
	}
}
