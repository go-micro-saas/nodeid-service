package biz

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	enumv1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/enums"
	errorv1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/errors"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/repo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
	datarepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/repo"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"github.com/rs/xid"
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
	var (
		now             = time.Now()
		initSerialModel = &po.NodeSerial{
			Id:            0,
			CreatedTime:   now,
			UpdatedTime:   now,
			InstanceId:    param.InstanceId,
			CurrentNodeId: 0,
		}
	)

	serialModel, err := s.nodeSerialData.FirstOrCreate(ctx, initSerialModel)
	if err != nil {
		return dataModel, err
	}

	// 线程
	tx := s.nodeSerialData.NewTransaction(ctx)
	defer func() {
		commitErr := tx.CommitAndErrRollback(ctx, err)
		if commitErr != nil {
			s.log.WithContext(ctx).Errorw(
				"mgs", "GetNodeId tx.CommitAndErrRollback failed!",
				"err", commitErr,
			)
		}
	}()
	// 锁行：排队获取；覆盖变量：serialModel
	serialModel, err = s.nodeSerialData.QueryOneByIdForUpdate(ctx, tx, serialModel.Id)
	if err != nil {
		return dataModel, err
	}
	// 获取一个已释放的闲置ID
	dataModel, isNotFound, err := s.nodeIDData.QueryOneIdleNodeIdByInstanceId(ctx, serialModel.InstanceId)
	if err != nil {
		return dataModel, err
	}
	if !isNotFound && dataModel != nil {
		return dataModel, err
	}
	// 获取一个已过期的ID
	expiredTime := s.conf.PreviousExpiredTime(now)
	dataModel, isNotFound, err = s.nodeIDData.QueryOneExpiredNodeIdByInstanceId(ctx, serialModel.InstanceId, expiredTime)
	if err != nil {
		return dataModel, err
	}
	if !isNotFound && dataModel != nil {
		s.renewalNodeIdData(dataModel)
		dataModel.AccessToken = s.genAccessToken()
		if err = s.nodeIDData.RenewalNodeIDWithTransaction(ctx, tx, dataModel); err != nil {
			return dataModel, err
		}
		return dataModel, err
	}
	// 实例化一个ID
	dataModel, err = s.GenerateNextID(serialModel, param)
	if err != nil {
		s.renewalNodeIdData(dataModel)
		dataModel.AccessToken = s.genAccessToken()
		if err = s.nodeIDData.RenewalNodeIDWithTransaction(ctx, tx, dataModel); err != nil {
			return dataModel, err
		}
		return dataModel, err
	}
	// 创建
	err = s.nodeIDData.CreateWithTransaction(ctx, tx, dataModel)
	if err != nil {
		return dataModel, err
	}
	// 更新当前ID
	serialModel.CurrentNodeId = dataModel.NodeId
	serialModel.UpdatedTime = now
	err = s.nodeSerialData.UpdateNodeIDWithTransaction(ctx, tx, serialModel)
	if err != nil {
		return dataModel, err
	}
	return dataModel, err
}

func (s *nodeIDBiz) GenerateNextID(serialModel *po.NodeSerial, param *bo.GetNodeIdParam) (*po.NodeId, error) {
	var nextID = serialModel.CurrentNodeId + 1
	if !s.conf.IsValidNodeID(nextID) {
		e := errorv1.DefaultErrorS102NoAvailableId()
		return nil, errorpkg.WithStack(e)
	}

	var (
		now = time.Now()
	)
	dataModel := &po.NodeId{
		Id:               0,
		CreatedTime:      now,
		UpdatedTime:      now,
		InstanceName:     param.InstanceName,
		InstanceId:       param.InstanceName,
		NodeId:           nextID,
		NodeIdStatus:     enumv1.NodeIDStatusEnum_USING,
		InstanceMetadata: nil,
		ExpiredAt:        s.conf.NextExpireTime(now),
		AccessToken:      s.genAccessToken(),
	}
	if param.Metadata == nil {
		param.Metadata = make(map[string]string)
	}
	metadataBuf, err := json.Marshal(param.Metadata)
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return dataModel, errorpkg.Wrap(e, err)
	}
	dataModel.InstanceMetadata = metadataBuf
	return dataModel, nil
}

func (s *nodeIDBiz) genAccessToken() string {
	return xid.New().String()
}

func (s *nodeIDBiz) RenewalNodeId(ctx context.Context, param *bo.RenewalNodeIdParam) (*po.NodeId, error) {
	dataModel, isNotFound, err := s.nodeIDData.QueryOneByInstanceNodeID(ctx, param.InstanceId, param.NodeID)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.DefaultErrorS102RecordNotFount()
		return nil, errorpkg.WithStack(e)
	}
	err = s.checkCanUpdateNodeID(dataModel, param.AccessToken)
	if err != nil {
		return nil, err
	}
	// 更新
	s.renewalNodeIdData(dataModel)
	err = s.nodeIDData.RenewalNodeID(ctx, dataModel)
	if err != nil {
		return nil, err
	}
	return dataModel, nil
}

func (s *nodeIDBiz) renewalNodeIdData(dataModel *po.NodeId) {
	now := time.Now()
	dataModel.UpdatedTime = now
	dataModel.NodeIdStatus = enumv1.NodeIDStatusEnum_USING
	dataModel.ExpiredAt = s.conf.NextExpireTime(now)
}

func (s *nodeIDBiz) ReleaseNodeId(ctx context.Context, param *bo.ReleaseNodeIdParam) (*po.NodeId, error) {
	dataModel, isNotFound, err := s.nodeIDData.QueryOneByInstanceNodeID(ctx, param.InstanceId, param.NodeID)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.DefaultErrorS102RecordNotFount()
		return nil, errorpkg.WithStack(e)
	}
	err = s.checkCanUpdateNodeID(dataModel, param.AccessToken)
	if err != nil {
		return nil, err
	}
	// 更新
	s.releaseNodeIdData(dataModel)
	err = s.nodeIDData.ReleaseNodeID(ctx, dataModel)
	if err != nil {
		return nil, err
	}
	return dataModel, nil
}

func (s *nodeIDBiz) releaseNodeIdData(dataModel *po.NodeId) {
	now := time.Now()
	dataModel.UpdatedTime = now
	dataModel.NodeIdStatus = enumv1.NodeIDStatusEnum_IDLE
	dataModel.ExpiredAt = now
}

func (s *nodeIDBiz) checkCanUpdateNodeID(dataModel *po.NodeId, accessToken string) error {
	if dataModel.AccessToken != accessToken {
		e := errorv1.DefaultErrorS102AccessTokenIncorrect()
		errorpkg.Kvs(e, "cause", "access token incorrect")
		return errorpkg.WithStack(e)
	}
	return nil
}
