package mq

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"gitlab.silkrode.com.tw/golang/mq/inside/pub"
	"gitlab.silkrode.com.tw/golang/mq/inside/sub"
)

// pub/sub message
type MQ struct {
	subscriber *sub.Subscriber
	publisher  *pub.Publisher
	ctx        context.Context
	cancel     context.CancelFunc
	client     *pubsub.Client
	logger     Logger
	Topic      *Topic
}

func NewMq(ctx context.Context) (*MQ, error) {
	client, err := NewClient(ctx)
	ctx, cacnel := context.WithCancel(ctx)
	if err != nil {
		return nil, err
	}
	return &MQ{
		ctx:    ctx,
		client: client,
		logger: &defaultLogger{},
		cancel: cacnel,
	}, nil
}

type Logger interface {
	InfoMsg(ctx context.Context, msg string, v ...interface{})
	ErrMsg(ctx context.Context, msg string, v ...interface{})
}

type defaultLogger struct{}

func (d *defaultLogger) InfoMsg(ctx context.Context, msg string, v ...interface{}) {
	log.Printf(msg, v)
}

func (d *defaultLogger) ErrMsg(ctx context.Context, msg string, v ...interface{}) {
	log.Printf(msg, v)
}

func (mq *MQ) SetLogger(logger Logger) {
	mq.logger = logger
}

func (mq *MQ) InitPublisher(topic ...string) error {
	if mq.publisher == nil {
		t := mq.Topic.topicID
		if len(topic) > 0 {
			t = topic[0]
		}
		mq.publisher = pub.NewPublish(mq.ctx, mq.client.Topic(t))
	}
	return nil
}

func (mq *MQ) GetPublisher() *pub.Publisher {
	return mq.publisher
}

// Publisher is obsolete. Use InitPublisher and GetPublisher
func (mq *MQ) Publisher() *pub.Publisher {
	mq.InitPublisher(mq.Topic.topicID)
	return mq.GetPublisher()
}

func (mq *MQ) InitSubscriber(subID ...string) error {
	if mq.subscriber == nil {
		s := ""
		if len(mq.Topic.gcpSubIDs) > 0 {
			s = mq.Topic.gcpSubIDs[0]
		} else if len(subID) > 0 {
			s = subID[0]
		}
		mq.subscriber = sub.NewSubscriber(mq.ctx, mq.client.Subscription(s))
	}
	return nil
}

func (mq *MQ) GetSubscriber() *sub.Subscriber {
	return mq.subscriber
}

// Publisher is obsolete. Use InitSubscriber and GetSubscriber
func (mq *MQ) Subscriber(subID ...string) *sub.Subscriber {
	mq.InitSubscriber(subID...)
	return mq.GetSubscriber()
}

func (mq *MQ) DeleteSubscription(ctx context.Context, subID string) error {
	sub := mq.client.Subscription(subID)
	if err := sub.Delete(ctx); err != nil {
		return fmt.Errorf("Delete sub: %v", err)
	}
	return nil
}

func (mq *MQ) DeleteAllSubscriptions(ctx context.Context) error {
	for _, subID := range mq.Topic.gcpSubIDs {
		if err := mq.DeleteSubscription(ctx, subID); err != nil {
			return err
		}
	}
	return nil
}

func (mq *MQ) Stop() {
	mq.cancel()
}
