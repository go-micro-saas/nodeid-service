// Package daos
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package daos

import (
	"bytes"
	context "context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	enumv1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/enums"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/repo"
	schemas "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/schema/node_id"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	gorm "gorm.io/gorm"
	"strings"
	"time"
)

// nodeIdRepo repo
type nodeIdRepo struct {
	log          *log.Helper
	dbConn       *gorm.DB       // *gorm.DB
	NodeIdSchema schemas.NodeId // NodeId
}

// NewNodeIdRepo new data repo
func NewNodeIdRepo(logger log.Logger, dbConn *gorm.DB) datarepos.NodeIdDataRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "nodeid-service/data/data"))
	return &nodeIdRepo{
		log:    logHelper,
		dbConn: dbConn,
	}
}

func (s *nodeIdRepo) NewTransaction(ctx context.Context, opts ...*sql.TxOptions) gormpkg.TransactionInterface {
	return gormpkg.NewTransaction(ctx, s.dbConn, opts...)
}

// =============== 创建 ===============

// create insert one
func (s *nodeIdRepo) create(ctx context.Context, dbConn *gorm.DB, dataModel *po.NodeId) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.NodeIdSchema.TableName()).
		Create(dataModel).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// Create insert one
func (s *nodeIdRepo) Create(ctx context.Context, dataModel *po.NodeId) error {
	return s.create(ctx, s.dbConn, dataModel)
}

// CreateWithDBConn create
func (s *nodeIdRepo) CreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.NodeId) error {
	return s.create(ctx, dbConn, dataModel)
}

// existCreate exist create
func (s *nodeIdRepo) existCreate(ctx context.Context, dbConn *gorm.DB, dataModel *po.NodeId) (anotherModel *po.NodeId, isNotFound bool, err error) {
	anotherModel = new(po.NodeId)
	err = dbConn.WithContext(ctx).
		Table(s.NodeIdSchema.TableName()).
		Where(schemas.FieldId+" = ?", dataModel.Id).
		First(anotherModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			isNotFound = true
			err = nil
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// ExistCreate exist create
func (s *nodeIdRepo) ExistCreate(ctx context.Context, dataModel *po.NodeId) (anotherModel *po.NodeId, isNotFound bool, err error) {
	return s.existCreate(ctx, s.dbConn, dataModel)
}

// ExistCreateWithDBConn exist create
func (s *nodeIdRepo) ExistCreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.NodeId) (anotherModel *po.NodeId, isNotFound bool, err error) {
	return s.existCreate(ctx, dbConn, dataModel)
}

// createInBatches create many
func (s *nodeIdRepo) createInBatches(ctx context.Context, dbConn *gorm.DB, dataModels []*po.NodeId, batchSize int) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.NodeIdSchema.TableName()).
		CreateInBatches(dataModels, batchSize).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// CreateInBatches create many
func (s *nodeIdRepo) CreateInBatches(ctx context.Context, dataModels []*po.NodeId, batchSize int) error {
	return s.createInBatches(ctx, s.dbConn, dataModels, batchSize)
}

// CreateInBatchesWithDBConn create many
func (s *nodeIdRepo) CreateInBatchesWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*po.NodeId, batchSize int) error {
	return s.createInBatches(ctx, dbConn, dataModels, batchSize)
}

// =============== 更新 ===============

// update update
func (s *nodeIdRepo) update(ctx context.Context, dbConn *gorm.DB, dataModel *po.NodeId) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.NodeIdSchema.TableName()).
		// Where(schemas.FieldId+" = ?", dataModel.Id).
		Save(dataModel).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// Update update
func (s *nodeIdRepo) Update(ctx context.Context, dataModel *po.NodeId) error {
	return s.update(ctx, s.dbConn, dataModel)
}

// UpdateWithDBConn update
func (s *nodeIdRepo) UpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.NodeId) error {
	return s.update(ctx, dbConn, dataModel)
}

// existUpdate exist update
func (s *nodeIdRepo) existUpdate(ctx context.Context, dbConn *gorm.DB, dataModel *po.NodeId) (anotherModel *po.NodeId, isNotFound bool, err error) {
	anotherModel = new(po.NodeId)
	err = dbConn.WithContext(ctx).
		Table(s.NodeIdSchema.TableName()).
		Where(schemas.FieldId+" = ?", dataModel.Id).
		Where(schemas.FieldId+" != ?", dataModel.Id).
		First(anotherModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			isNotFound = true
			err = nil
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// ExistUpdate exist update
func (s *nodeIdRepo) ExistUpdate(ctx context.Context, dataModel *po.NodeId) (anotherModel *po.NodeId, isNotFound bool, err error) {
	return s.existUpdate(ctx, s.dbConn, dataModel)
}

// ExistUpdateWithDBConn exist update
func (s *nodeIdRepo) ExistUpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.NodeId) (anotherModel *po.NodeId, isNotFound bool, err error) {
	return s.existUpdate(ctx, dbConn, dataModel)
}

// =============== query one : 查一个 ===============

// queryOneById query one by id
func (s *nodeIdRepo) queryOneById(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *po.NodeId, isNotFound bool, err error) {
	dataModel = new(po.NodeId)
	err = dbConn.WithContext(ctx).
		Table(s.NodeIdSchema.TableName()).
		Where(schemas.FieldId+" = ?", id).
		First(dataModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// QueryOneById query one by id
func (s *nodeIdRepo) QueryOneById(ctx context.Context, id interface{}) (dataModel *po.NodeId, isNotFound bool, err error) {
	return s.queryOneById(ctx, s.dbConn, id)
}

// QueryOneByIdWithDBConn query one by id
func (s *nodeIdRepo) QueryOneByIdWithDBConn(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *po.NodeId, isNotFound bool, err error) {
	return s.queryOneById(ctx, dbConn, id)
}

func (s *nodeIdRepo) QueryOneIdleNodeIdByInstanceId(ctx context.Context, instanceID string) (dataModel *po.NodeId, isNotFound bool, err error) {
	dataModel = new(po.NodeId)
	err = s.dbConn.WithContext(ctx).
		Table(s.NodeIdSchema.TableName()).
		Where(schemas.FieldInstanceId+" = ?", instanceID).
		Where(schemas.FieldNodeIdStatus+" = ?", enumv1.NodeIDStatusEnum_IDLE).
		Order(schemas.FieldId).
		First(dataModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

func (s *nodeIdRepo) QueryOneExpiredNodeIdByInstanceId(ctx context.Context, instanceID string, expiredTime time.Time) (dataModel *po.NodeId, isNotFound bool, err error) {
	dataModel = new(po.NodeId)
	err = s.dbConn.WithContext(ctx).
		Table(s.NodeIdSchema.TableName()).
		Where(schemas.FieldInstanceId+" = ?", instanceID).
		Where(schemas.FieldExpiredAt+" <= ?", expiredTime).
		Order(schemas.FieldExpiredAt).
		First(dataModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// queryOneByConditions query one by conditions
func (s *nodeIdRepo) queryOneByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *po.NodeId, isNotFound bool, err error) {
	dataModel = new(po.NodeId)
	dbConn = dbConn.WithContext(ctx).Table(s.NodeIdSchema.TableName())
	err = s.WhereConditions(dbConn, conditions).
		First(dataModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// QueryOneByConditions query one by conditions
func (s *nodeIdRepo) QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *po.NodeId, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, s.dbConn, conditions)
}

// QueryOneByConditionsWithDBConn query one by conditions
func (s *nodeIdRepo) QueryOneByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *po.NodeId, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, dbConn, conditions)
}

// =============== query all : 查全部 ===============

// queryAllByConditions query all by conditions
func (s *nodeIdRepo) queryAllByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModels []*po.NodeId, err error) {
	dbConn = dbConn.WithContext(ctx).Table(s.NodeIdSchema.TableName())
	err = s.WhereConditions(dbConn, conditions).
		Find(&dataModels).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return dataModels, err
	}
	return
}

// QueryAllByConditions query all by conditions
func (s *nodeIdRepo) QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) ([]*po.NodeId, error) {
	return s.queryAllByConditions(ctx, s.dbConn, conditions)
}

// QueryAllByConditionsWithDBConn query all by conditions
func (s *nodeIdRepo) QueryAllByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) ([]*po.NodeId, error) {
	return s.queryAllByConditions(ctx, dbConn, conditions)
}

// =============== list : 列表 ===============

// list 列表
func (s *nodeIdRepo) list(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) (dataModels []*po.NodeId, recordCount int64, err error) {
	// query where
	dbConn = dbConn.WithContext(ctx).Table(s.NodeIdSchema.TableName())
	dbConn = s.WhereConditions(dbConn, conditions)
	dbConn = gormpkg.AssembleWheres(dbConn, paginatorArgs.PageWheres)

	err = dbConn.Count(&recordCount).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return
	} else if recordCount == 0 {
		return // empty
	}

	// pagination
	dbConn = gormpkg.AssembleOrders(dbConn, paginatorArgs.PageOrders)
	err = gormpkg.Paginator(dbConn, paginatorArgs.PageOption).
		Find(&dataModels).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return
	}
	return
}

// List 列表
func (s *nodeIdRepo) List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) ([]*po.NodeId, int64, error) {
	return s.list(ctx, s.dbConn, conditions, paginatorArgs)
}

// ListWithDBConn 列表
func (s *nodeIdRepo) ListWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) ([]*po.NodeId, int64, error) {
	return s.list(ctx, dbConn, conditions, paginatorArgs)
}

// =============== delete : 删除 ===============

// delete delete one
func (s *nodeIdRepo) delete(ctx context.Context, dbConn *gorm.DB, dataModel *po.NodeId) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.NodeIdSchema.TableName()).
		Where(schemas.FieldId+" = ?", dataModel.Id).
		Delete(dataModel).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return
}

// Delete delete one
func (s *nodeIdRepo) Delete(ctx context.Context, dataModel *po.NodeId) error {
	return s.delete(ctx, s.dbConn, dataModel)
}

// DeleteWithDBConn delete one
func (s *nodeIdRepo) DeleteWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.NodeId) error {
	return s.delete(ctx, dbConn, dataModel)
}

// deleteByIds delete by ids
func (s *nodeIdRepo) deleteByIds(ctx context.Context, dbConn *gorm.DB, ids interface{}) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.NodeIdSchema.TableName()).
		Where(schemas.FieldId+" in (?)", ids).
		Delete(po.NodeId{}).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return
}

// DeleteByIds delete by ids
func (s *nodeIdRepo) DeleteByIds(ctx context.Context, ids interface{}) error {
	return s.deleteByIds(ctx, s.dbConn, ids)
}

// DeleteByIdsWithDBConn delete by ids
func (s *nodeIdRepo) DeleteByIdsWithDBConn(ctx context.Context, dbConn *gorm.DB, ids interface{}) error {
	return s.deleteByIds(ctx, dbConn, ids)
}

// =============== insert : 批量入库 ===============

var _ gormpkg.BatchInsertRepo = new(NodeIdSlice)

// NodeIdSlice 表切片
type NodeIdSlice []*po.NodeId

// TableName 表名
func (s *NodeIdSlice) TableName() string {
	return schemas.NodeIdSchema.TableName()
}

// Len 长度
func (s *NodeIdSlice) Len() int {
	return len(*s)
}

// InsertColumns 批量入库的列
func (s *NodeIdSlice) InsertColumns() (columnList []string, placeholder string) {
	// columns
	columnList = []string{
		schemas.FieldCreatedTime, schemas.FieldUpdatedTime,
		schemas.FieldInstanceName, schemas.FieldInstanceId,
		schemas.FieldNodeId, schemas.FieldNodeIdStatus,
		schemas.FieldInstanceMetadata, schemas.FieldExpiredAt,
	}

	// placeholders
	insertPlaceholderBytes := bytes.Repeat([]byte("?, "), len(columnList))
	insertPlaceholderBytes = bytes.TrimSuffix(insertPlaceholderBytes, []byte(", "))

	return columnList, string(insertPlaceholderBytes)
}

// InsertValues 批量入库的值
func (s *NodeIdSlice) InsertValues(args *gormpkg.BatchInsertValueArgs) (prepareData []interface{}, placeholderSlice []string) {
	dataModels := (*s)[args.StepStart:args.StepEnd]
	for index := range dataModels {
		// placeholder
		placeholderSlice = append(placeholderSlice, "("+args.InsertPlaceholder+")")

		// prepare data
		prepareData = append(prepareData, dataModels[index].CreatedTime)
		prepareData = append(prepareData, dataModels[index].UpdatedTime)
		prepareData = append(prepareData, dataModels[index].InstanceName)
		prepareData = append(prepareData, dataModels[index].InstanceId)
		prepareData = append(prepareData, dataModels[index].NodeId)
		prepareData = append(prepareData, dataModels[index].NodeIdStatus)
		prepareData = append(prepareData, dataModels[index].InstanceMetadata)
		prepareData = append(prepareData, dataModels[index].ExpiredAt)
	}
	return prepareData, placeholderSlice
}

// UpdateColumns 批量入库的列
func (s *NodeIdSlice) UpdateColumns() (columnList []string) {
	// columns
	columnList = []string{
		schemas.FieldCreatedTime + "=excluded." + schemas.FieldCreatedTime,
		schemas.FieldUpdatedTime + "=excluded." + schemas.FieldUpdatedTime,
		schemas.FieldInstanceName + "=excluded." + schemas.FieldInstanceName,
		schemas.FieldInstanceId + "=excluded." + schemas.FieldInstanceId,
		schemas.FieldNodeId + "=excluded." + schemas.FieldNodeId,
		schemas.FieldNodeIdStatus + "=excluded." + schemas.FieldNodeIdStatus,
		schemas.FieldInstanceMetadata + "=excluded." + schemas.FieldInstanceMetadata,
		schemas.FieldExpiredAt + "=excluded." + schemas.FieldExpiredAt,
	}
	return columnList
}

// ConflictActionForMySQL 更新冲突时的操作
func (s *NodeIdSlice) ConflictActionForMySQL() (req *gormpkg.BatchInsertConflictActionReq) {
	req = &gormpkg.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "AS excluded",
		OnConflictTarget:      "ON DUPLICATE KEY",
		OnConflictAction:      "UPDATE " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// ConflictActionForPostgres 更新冲突时的操作
func (s *NodeIdSlice) ConflictActionForPostgres() (req *gormpkg.BatchInsertConflictActionReq) {
	req = &gormpkg.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "",
		OnConflictTarget:      "ON CONFLICT(id)",
		OnConflictAction:      "DO UPDATE SET " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// insert 批量插入
func (s *nodeIdRepo) insert(ctx context.Context, dbConn *gorm.DB, dataModels NodeIdSlice) error {
	err := gormpkg.BatchInsertWithContext(ctx, dbConn, &dataModels)
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return nil
}

// Insert 批量插入
func (s *nodeIdRepo) Insert(ctx context.Context, dataModels []*po.NodeId) error {
	return s.insert(ctx, s.dbConn, dataModels)
}

// InsertWithDBConn 批量插入
func (s *nodeIdRepo) InsertWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*po.NodeId) error {
	return s.insert(ctx, dbConn, dataModels)
}

// =============== conditions : 条件 ===============

// WhereConditions orm where
func (s *nodeIdRepo) WhereConditions(dbConn *gorm.DB, conditions map[string]interface{}) *gorm.DB {

	// table name
	// tableName := s.NodeIdSchema.TableName()

	// On-demand loading

	// id
	// if data, ok := conditions[schemas.FieldId]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldId+" = ?", data)
	// }

	// created_time
	// if data, ok := conditions[schemas.FieldCreatedTime]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldCreatedTime+" = ?", data)
	// }

	// updated_time
	// if data, ok := conditions[schemas.FieldUpdatedTime]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldUpdatedTime+" = ?", data)
	// }

	// instance_name
	// if data, ok := conditions[schemas.FieldInstanceName]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldInstanceName+" = ?", data)
	// }

	// instance_id
	// if data, ok := conditions[schemas.FieldInstanceId]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldInstanceId+" = ?", data)
	// }

	// node_id
	// if data, ok := conditions[schemas.FieldNodeId]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldNodeId+" = ?", data)
	// }

	// node_id_status
	// if data, ok := conditions[schemas.FieldNodeIdStatus]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldNodeIdStatus+" = ?", data)
	// }

	// instance_metadata
	// if data, ok := conditions[schemas.FieldInstanceMetadata]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldInstanceMetadata+" = ?", data)
	// }

	// expired_at
	// if data, ok := conditions[schemas.FieldExpiredAt]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldExpiredAt+" = ?", data)
	// }

	return dbConn
}

// UpdateColumns update columns
func (s *nodeIdRepo) UpdateColumns(conditions map[string]interface{}) map[string]interface{} {

	// update columns
	updateColumns := make(map[string]interface{})

	// On-demand loading

	// id
	//if data, ok := conditions[schemas.FieldId]; ok {
	//	updateColumns[schemas.FieldId] = data
	//}

	// created_time
	//if data, ok := conditions[schemas.FieldCreatedTime]; ok {
	//	updateColumns[schemas.FieldCreatedTime] = data
	//}

	// updated_time
	//if data, ok := conditions[schemas.FieldUpdatedTime]; ok {
	//	updateColumns[schemas.FieldUpdatedTime] = data
	//}

	// instance_name
	//if data, ok := conditions[schemas.FieldInstanceName]; ok {
	//	updateColumns[schemas.FieldInstanceName] = data
	//}

	// instance_id
	//if data, ok := conditions[schemas.FieldInstanceId]; ok {
	//	updateColumns[schemas.FieldInstanceId] = data
	//}

	// node_id
	//if data, ok := conditions[schemas.FieldNodeId]; ok {
	//	updateColumns[schemas.FieldNodeId] = data
	//}

	// node_id_status
	//if data, ok := conditions[schemas.FieldNodeIdStatus]; ok {
	//	updateColumns[schemas.FieldNodeIdStatus] = data
	//}

	// instance_metadata
	//if data, ok := conditions[schemas.FieldInstanceMetadata]; ok {
	//	updateColumns[schemas.FieldInstanceMetadata] = data
	//}

	// expired_at
	//if data, ok := conditions[schemas.FieldExpiredAt]; ok {
	//	updateColumns[schemas.FieldExpiredAt] = data
	//}

	return updateColumns
}
