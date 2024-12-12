package biz

import (
	"context"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
	"testing"
	"time"
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

// go test -v -count=1 ./app/nodeid-service/internal/biz/biz -test.run=Test_nodeIDBiz_RenewalNodeId
func Test_nodeIDBiz_RenewalNodeId(t *testing.T) {

	type args struct {
		ctx   context.Context
		param *bo.RenewalNodeIdParam
	}
	tests := []struct {
		name    string
		args    args
		want    *po.NodeId
		wantErr bool
	}{
		{
			name: "#RenewalNodeId",
			args: args{
				ctx: context.Background(),
				param: &bo.RenewalNodeIdParam{
					InstanceId:  "testdata",
					NodeID:      2,
					AccessToken: "",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nodeIDBizHandler.RenewalNodeId(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenewalNodeId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("RenewalNodeId() got = %v, want %v", got, tt.want)
			//}
			t.Logf("==> got: %#v\n", got)
			t.Logf("==> got.ExpiredAt: %v\n", got.ExpiredAt)
			t.Logf("==> got.UpdatedTime: %v\n", got.UpdatedTime)
			t.Logf("==> now: %v\n", time.Now())
			t.Logf("==> nodeIDBizHandler.conf.IdleDuration: %v\n", nodeIDBizHandler.conf.IdleDuration)
			t.Logf("==> got.ExpiredAt.Sub(got.UpdatedTime): %v\n", got.ExpiredAt.Sub(got.UpdatedTime))
			t.Logf("==> got.ExpiredAt.Sub(time.Now()): %v\n", got.ExpiredAt.Sub(time.Now()))
		})
	}
}

// go test -v -count=1 ./app/nodeid-service/internal/biz/biz -test.run=Test_nodeIDBiz_ReleaseNodeId
func Test_nodeIDBiz_ReleaseNodeId(t *testing.T) {

	type args struct {
		ctx   context.Context
		param *bo.ReleaseNodeIdParam
	}
	tests := []struct {
		name    string
		args    args
		want    *po.NodeId
		wantErr bool
	}{
		{
			name: "#RenewalNodeId",
			args: args{
				ctx: context.Background(),
				param: &bo.ReleaseNodeIdParam{
					InstanceId:  "testdata",
					NodeID:      2,
					AccessToken: "",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nodeIDBizHandler.ReleaseNodeId(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReleaseNodeId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("RenewalNodeId() got = %v, want %v", got, tt.want)
			//}
			t.Logf("==> got: %#v\n", got)
			t.Logf("==> got.NodeIdStatus: %v\n", got.NodeIdStatus)
			t.Logf("==> got.ExpiredAt: %v\n", got.ExpiredAt)
			t.Logf("==> got.UpdatedTime: %v\n", got.UpdatedTime)
		})
	}
}
