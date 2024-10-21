package bizrepos

import "context"

type UuidBizRepo interface {
	NextID(ctx context.Context) (uint64, error)
}
