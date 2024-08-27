package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/repo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
	datarepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/repo"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"time"
)

type nodeIDBiz struct {
	log            *log.Helper
	conf           *bo.NodeIDConfig
	nodeIDData     datarepos.NodeIdDataRepo
	nodeSerialData datarepos.NodeSerialDataRepo
}

func NewNodeIDBiz(
	logger log.Logger,
	conf *bo.NodeIDConfig,
	nodeIDData datarepos.NodeIdDataRepo,
	nodeSerialData datarepos.NodeSerialDataRepo,
) bizrepos.NodeIdBizRepo {
	conf.Initialization()

	logHelper := log.NewHelper(log.With(logger, "module", "nodeid-service/biz/biz"))
	return &nodeIDBiz{
		log:            logHelper,
		conf:           conf,
		nodeIDData:     nodeIDData,
		nodeSerialData: nodeSerialData,
	}
}

func (s *nodeIDBiz) GetConfig() *bo.NodeIDConfig {
	return s.conf.Clone()
}

func (s *nodeIDBiz) GetNodeId(ctx context.Context, param *bo.GetNodeIdParam) (dataModel *po.NodeId, err error) {
	initSerialModel := &po.NodeSerial{
		Id:            0,
		CreatedTime:   time.Now(),
		UpdatedTime:   time.Now(),
		InstanceId:    param.InstanceId,
		CurrentNodeId: 0,
	}
	serialModel, err := s.nodeSerialData.FirstOrCreate(ctx, initSerialModel)
	if err != nil {
		return dataModel, err
	}

	// 锁行：排队获取
	tx := s.nodeSerialData.NewTransaction(ctx)
	defer func() {
		commitErr := tx.CommitAndErrRollback(ctx, err)
		if commitErr != nil {
			e := errorpkg.ErrorInternalError("")
			err = errorpkg.Wrap(e, commitErr)
		}
	}()
	serialModel, err = s.nodeSerialData.QueryOneByIdForUpdate(ctx, tx, serialModel.Id)
	if err != nil {
		return dataModel, err
	}
	return dataModel, err
}
