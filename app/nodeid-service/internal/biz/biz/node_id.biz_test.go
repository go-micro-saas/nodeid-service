package biz

import (
	"context"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
	"testing"
)

// go test -v -count=1 ./app/nodeid-service/internal/biz/biz -test.run=Test_nodeIDBiz_GetNodeId
func Test_nodeIDBiz_GetNodeId(t *testing.T) {
	type args struct {
		ctx   context.Context
		param *bo.GetNodeIdParam
	}
	tests := []struct {
		name          string
		args          args
		wantDataModel *po.NodeId
		wantErr       bool
	}{
		{
			name: "#GetNodeId",
			args: args{
				ctx: context.Background(),
				param: &bo.GetNodeIdParam{
					InstanceId:   "testdata",
					InstanceName: "testdata",
					Metadata:     map[string]string{"k": "v"},
				},
			},
			wantDataModel: nil,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDataModel, err := nodeIDBizHandler.GetNodeId(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNodeId() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotDataModel, tt.wantDataModel) {
			//	t.Errorf("GetNodeId() gotDataModel = %v, want %v", gotDataModel, tt.wantDataModel)
			//}
			t.Logf("==> got: %#v\n", gotDataModel)
		})
	}
}
