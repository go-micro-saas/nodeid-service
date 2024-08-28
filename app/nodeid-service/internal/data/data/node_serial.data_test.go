package data

import (
	"context"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	"testing"
	"time"
)

// go test -v -count=1 ./app/nodeid-service/internal/data/data -test.run=Test_nodeSerialRepo_IsErrDuplicatedKey_Ignore
func Test_nodeSerialRepo_IsErrDuplicatedKey_Ignore(t *testing.T) {
	dataModel := &po.NodeSerial{
		Id:            0,
		CreatedTime:   time.Now(),
		UpdatedTime:   time.Now(),
		InstanceId:    "testdata",
		CurrentNodeId: 0,
	}
	type args struct {
		ctx       context.Context
		dataModel *po.NodeSerial
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#create_IsErrDuplicatedKey_Ignore",
			args: args{
				ctx:       context.Background(),
				dataModel: dataModel,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 已被转换，不能判断 DuplicatedKey
			//err := nodeSerialHandler.Create(tt.args.ctx, tt.args.dataModel)
			// 原生的未被转换
			err := dbConnection.Table(nodeSerialHandler.NodeSerialSchema.TableName()).Create(tt.args.dataModel).Error
			if err != nil && gormpkg.IsErrDuplicatedKey(err) {
				t.Log("==> IsErrDuplicatedKey: ", err)
				err = nil
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// go test -v -count=1 ./app/nodeid-service/internal/data/data -test.run=Test_nodeSerialRepo_FirstOrCreate
func Test_nodeSerialRepo_FirstOrCreate(t *testing.T) {
	dataModel := &po.NodeSerial{
		Id:            0,
		CreatedTime:   time.Now(),
		UpdatedTime:   time.Now(),
		InstanceId:    "testing",
		CurrentNodeId: 0,
	}
	type args struct {
		ctx       context.Context
		dataModel *po.NodeSerial
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#FirstOrCreate",
			args: args{
				ctx:       context.Background(),
				dataModel: dataModel,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nodeSerialHandler.FirstOrCreate(tt.args.ctx, tt.args.dataModel)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstOrCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("==> got: %#v\n", got)
		})
	}
}
