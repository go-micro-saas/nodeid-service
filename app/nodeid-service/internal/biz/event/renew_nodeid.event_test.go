package events

import (
	"context"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/repo"
	threadpkg "github.com/ikaiguang/go-srv-kit/kratos/thread"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// go test -v -count=1 ./app/nodeid-service/internal/biz/event -test.run=Test_renewNodeIDEvent_Receive
func Test_renewNodeIDEvent_Receive(t *testing.T) {
	fakeHandler := func(ctx context.Context, param *bo.RenewalNodeIdParam) (*bo.RenewalNodeIDReply, error) {
		return &bo.RenewalNodeIDReply{}, nil
	}
	type args struct {
		ctx            context.Context
		handler        bizrepos.RenewNodeIDHandler
		sendMessageNum int
	}
	sendMessageNum := 3
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#Test_renewNodeIDEvent_Receive",
			args: args{
				ctx:            context.Background(),
				handler:        fakeHandler,
				sendMessageNum: sendMessageNum,
			},
			wantErr: false,
		},
	}

	threadpkg.GoSafe(func() {
		t.Logf("==> start renewNodeIDEvent_Receive\n")
		defer func() { t.Logf("==> end renewNodeIDEvent_Receive\n") }()
		err := handler.Consume(context.Background(), fakeHandler)
		require.Nil(t, err)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if err := handler.Consume(tt.args.ctx, tt.args.handler); (err != nil) != tt.wantErr {
			//	t.Errorf("Consume() error = %v, wantErr %v", err, tt.wantErr)
			//}
			param := &bo.RenewalNodeIdParam{
				InstanceId:  "testdata",
				NodeID:      1,
				AccessToken: "ctikq3cfegt0l8m10rgg",
			}
			for i := 1; i <= tt.args.sendMessageNum; i++ {
				t.Logf("==> renewNodeIDEvent_Send; num:%d\n", i)
				err := handler.Publish(tt.args.ctx, param)
				require.Nil(t, err)
			}
		})
	}

	for {
		time.Sleep(time.Second)
		t.Logf("==> time.Second; receiveCounter:%d\n", handler.receiveCounter)
		if handler.receiveCounter >= uint64(sendMessageNum) {
			t.Logf("==> Test_renewNodeIDEvent_Receive : sendMessageNum : %d\n", sendMessageNum)
			err := handler.Close(context.Background())
			require.Nil(t, err)
			return
		}
	}
}

// go test -v -count=1 ./app/nodeid-service/internal/biz/event -test.run=Test_renewNodeIDEvent_Send
func Test_renewNodeIDEvent_Send(t *testing.T) {

	type args struct {
		ctx   context.Context
		param *bo.RenewalNodeIdParam
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#Test_renewNodeIDEvent_Send",
			args: args{
				ctx: context.Background(),
				param: &bo.RenewalNodeIdParam{
					InstanceId:  "testdata1",
					NodeID:      1,
					AccessToken: "xxx",
				},
			},
			wantErr: false,
		},
		{
			name: "#Test_renewNodeIDEvent_Send",
			args: args{
				ctx: context.Background(),
				param: &bo.RenewalNodeIdParam{
					InstanceId:  "testdata2",
					NodeID:      2,
					AccessToken: "xxx",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handler.Publish(tt.args.ctx, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
