package bizrepos

import (
	"context"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
)

type RenewNodeIDHandler func(ctx context.Context, param *bo.RenewalNodeIdParam) error

type RenewNodeIDEventRepo interface {
	Publish(ctx context.Context, param *bo.RenewalNodeIdParam) error
	Consume(ctx context.Context, handler RenewNodeIDHandler) error
	Handle(ctx context.Context, param *bo.RenewalNodeIdParam) error
	Close(ctx context.Context) error
}
