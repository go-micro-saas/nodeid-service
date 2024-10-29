package uuidapi

import (
	"context"
	nodeidresourcev1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
)

type UuidHelper interface {
	GetSnowflakeNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error)
}

func GetContext() context.Context {
	return context.Background()
}

func GetSnowflakeNode(ctx context.Context, helper UuidHelper, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error) {
	return helper.GetSnowflakeNode(ctx, req)
}
