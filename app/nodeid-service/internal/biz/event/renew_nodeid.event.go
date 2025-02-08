package events

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-kratos/kratos/v2/log"
	enumv1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/enums"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/biz/repo"
	"github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/po"
	datarepos "github.com/go-micro-saas/nodeid-service/app/nodeid-service/internal/data/repo"
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
	logger           log.Logger
	log              *log.Helper
	mqConn           *rabbitmqpkg.ConnectionWrapper
	eventHistoryRepo datarepos.NodeEventHistoryRepo

	topic          string             // topic
	initPubSubOnce sync.Once          // init
	pub            message.Publisher  // 使用 getPublisherSubscriber
	sub            message.Subscriber // 使用 getPublisherSubscriber

	closing        chan struct{}
	receiveCounter uint64
}

func NewRenewNodeIDEventRepo(
	logger log.Logger,
	mqConn *rabbitmqpkg.ConnectionWrapper,
	eventHistoryRepo datarepos.NodeEventHistoryRepo,
) bizrepos.RenewNodeIDEventRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "nodeid-service/biz/event/renew_nodeid"))
	return &renewNodeIDEvent{
		logger:           logger,
		log:              logHelper,
		mqConn:           mqConn,
		eventHistoryRepo: eventHistoryRepo,
		topic:            po.Key(RenewNodeIDEventTopic),
		closing:          make(chan struct{}),
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

func (s *renewNodeIDEvent) Publish(ctx context.Context, param *bo.RenewalNodeIdParam) error {
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
	err = publisher.Publish(s.topic, msg)
	if err != nil {
		return errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
	}
	return nil
}

func (s *renewNodeIDEvent) Consume(ctx context.Context, handler bizrepos.RenewNodeIDHandler) error {
	_, subscriber, err := s.getPublisherSubscriber()
	if err != nil {
		return err
	}
	m, err := subscriber.Subscribe(context.Background(), s.topic)
	if err != nil {
		return errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
	}
	for {
		select {
		case msg := <-m:
			s.receiveCounter++
			_ = s.processMessage(ctx, handler, msg)
		case <-s.closing:
			s.log.Debugw("msg", "RenewNodeIDEvent.Consume Stopping Consume")
			return nil
		}
	}
}

func (s *renewNodeIDEvent) Close(ctx context.Context) error {
	s.closing <- struct{}{}
	return nil
}

func (s *renewNodeIDEvent) processMessage(ctx context.Context, handler bizrepos.RenewNodeIDHandler, msg *message.Message) (err error) {
	defer func() {
		historyErr := s.saveHistory(ctx, msg, err)
		if historyErr != nil {
			s.log.WithContext(ctx).Errorw("msg", "RenewNodeIDEvent.Consume SaveHistory failed",
				"err", err, "historyErr", string(msg.Payload))
			return
		}
		msg.Ack()
	}()
	s.log.WithContext(ctx).Infow("msg", "RenewNodeIDEvent.Consume Received message", "receiveCounter", s.receiveCounter, "msg.payload", string(msg.Payload))
	param := &bo.RenewalNodeIdParam{}
	err = param.UnmarshalFromJSON(msg.Payload)
	if err != nil {
		s.log.WithContext(ctx).Errorw("msg", "RenewNodeIDEvent.Consume RenewalNodeIdParam UnmarshalFromJSON failed",
			"err", err, "payload", string(msg.Payload))
		return err
	}
	result, err := handler(context.Background(), param)
	if err != nil {
		s.log.WithContext(ctx).Errorw("msg", "RenewNodeIDEvent.Consume Process failed",
			"err", err, "payload", string(msg.Payload), "result", result)
		return err
	}
	return err
}

func (s *renewNodeIDEvent) Name() string {
	return "renewNodeIDEvent"
}

func (s *renewNodeIDEvent) saveHistory(ctx context.Context, processMsg *message.Message, processErr error) error {
	historyModel := po.DefaultNodeEventHistory(s.Name(), string(processMsg.Payload))
	if processErr != nil {
		historyModel.EventStatus = enumv1.NodeEventStatusEnum_FAILED
		historyModel.EventError = fmt.Sprintf("%+v", processErr)
	}
	err := s.eventHistoryRepo.Create(ctx, historyModel)
	if err != nil {
		return err
	}
	return nil
}
