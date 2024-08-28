package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
	datarepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/repo"
	"reflect"
	"testing"
)

func Test_nodeIDBiz_GetNodeId(t *testing.T) {
	type fields struct {
		log            *log.Helper
		conf           *bo.NodeIDConfig
		nodeIDData     datarepos.NodeIdDataRepo
		nodeSerialData datarepos.NodeSerialDataRepo
	}
	type args struct {
		ctx   context.Context
		param *bo.GetNodeIdParam
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantDataModel *po.NodeId
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &nodeIDBiz{
				log:            tt.fields.log,
				conf:           tt.fields.conf,
				nodeIDData:     tt.fields.nodeIDData,
				nodeSerialData: tt.fields.nodeSerialData,
			}
			gotDataModel, err := s.GetNodeId(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNodeId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDataModel, tt.wantDataModel) {
				t.Errorf("GetNodeId() gotDataModel = %v, want %v", gotDataModel, tt.wantDataModel)
			}
		})
	}
}
