package events

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/repo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
	rabbitmqpkg "github.com/ikaiguang/go-srv-kit/data/rabbitmq"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

var RenewNodeIDEventTopic = "renew_node_id"

var (
	_ bizrepos.RenewNodeIDEventRepo = (*renewNodeIDEvent)(nil)
)

type renewNodeIDEvent struct {
	logger        log.Logger
	log           *log.Helper
	mqConn        *rabbitmqpkg.ConnectionWrapper
	nodeidBizRepo bizrepos.NodeIdBizRepo

	topic          string             // topic
	initPubSubOnce sync.Once          // init
	pub            message.Publisher  // 使用 getPublisherSubscriber
	sub            message.Subscriber // 使用 getPublisherSubscriber

	ctx            context.Context // context
	closing        chan struct{}
	receiveCounter uint64
}

func NewRenewNodeIDEventRepo(
	logger log.Logger,
	mqConn *rabbitmqpkg.ConnectionWrapper,
	nodeidBizRepo bizrepos.NodeIdBizRepo,
) bizrepos.RenewNodeIDEventRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "nodeid-service/biz/event/renew_nodeid"))
	return &renewNodeIDEvent{
		logger:        logger,
		log:           logHelper,
		mqConn:        mqConn,
		nodeidBizRepo: nodeidBizRepo,
		topic:         po.Key(RenewNodeIDEventTopic),
		ctx:           context.Background(),
		closing:       make(chan struct{}),
	}
}

func (s *renewNodeIDEvent) getPublisherSubscriber() (message.Publisher, message.Subscriber, error) {
	var (
		err error
	)
	s.initPubSubOnce.Do(func() {
		s.pub, s.sub, err = rabbitmqpkg.NewPublisherAndSubscriberWithConnection(s.mqConn, rabbitmqpkg.WithKratosLogger(s.logger))
	})
	if err != nil {
		s.initPubSubOnce = sync.Once{}
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, nil, errorpkg.WithStack(e)
	}
	return s.pub, s.sub, nil
}

func (s *renewNodeIDEvent) Send(ctx context.Context, param *bo.RenewalNodeIdParam) error {
	publisher, _, err := s.getPublisherSubscriber()
	if err != nil {
		return err
	}
	payload, err := param.MarshalToJSON()
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuidpkg.New(), payload)
	msg.SetContext(ctx)
	err = publisher.Publish(s.topic)
	if err != nil {
		return errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
	}
	return nil
}

func (s *renewNodeIDEvent) Receive(ctx context.Context, handler bizrepos.RenewNodeIDHandler) error {
	_, subscriber, err := s.getPublisherSubscriber()
	if err != nil {
		return err
	}
	m, err := subscriber.Subscribe(s.ctx, s.topic)
	if err != nil {
		return errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
	}
	for {
		select {
		case msg := <-m:
			s.receiveCounter++
			{
				param := &bo.RenewalNodeIdParam{}
				err := param.UnmarshalFromJSON(msg.Payload)
				if err != nil {
					s.log.WithContext(ctx).Errorw("msg", "RenewNodeIDEvent.Receive RenewalNodeIdParam UnmarshalFromJSON failed", "err", err)
					continue
				}
				err = s.Process(s.ctx, param)
				if err != nil {
					s.log.WithContext(ctx).Errorw("msg", "RenewNodeIDEvent.Receive Process failed", "err", err)
					continue
				}
				msg.Ack()
			}
		case <-s.closing:
			s.log.Debugw("msg", "RenewNodeIDEvent.Receive Stopping Receive")
			return nil
		}
	}
	//return nil
}

func (s *renewNodeIDEvent) Process(ctx context.Context, param *bo.RenewalNodeIdParam) error {
	_, err := s.nodeidBizRepo.RenewalNodeId(ctx, param)
	if err != nil {
		return err
	}
	return nil
}

func (s *renewNodeIDEvent) Close(ctx context.Context) error {
	s.closing <- struct{}{}
	return nil
}
