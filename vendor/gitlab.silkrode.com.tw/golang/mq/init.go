package mq

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.silkrode.com.tw/golang/mq/inside/pub"
	"gitlab.silkrode.com.tw/golang/mq/inside/sub"
)

var (
	CredentialsFile string
	ProjectID       string
)

func Init(ctx context.Context, topicID string, credentialsFile string, projectID string, opts ...Option) (*MQ, error) {
	CredentialsFile = credentialsFile
	ProjectID = projectID
	t := &Topic{
		ctx:     ctx,
		initPub: false,
		initSub: false,
		topicID: topicID,
	}

	for _, opt := range opts {
		opt(t)
	}

	if t.wg != nil {
		defer t.wg.Done()
		t.wg.Add(1)
	}

	client, err := NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "client error")
	}

	//new mq instance
	mqInstance, err := NewMq(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "new mq error")
	}

	mqInstance.Topic = t

	//pub/sub 需要做的
	err = t.CreateTopicIfNotExist(client, t.topicID)
	if err != nil {
		return nil, errors.Wrap(err, "create topic not exist error")
	}

	for _, gcpSubID := range t.gcpSubIDs {
		//create subscription by subID
		err := t.CreateSubscriptionIfNotExist(client, gcpSubID)
		if err != nil {
			return nil, errors.Wrap(err, "sub create subscription if not exist error")
		}
	}

	if t.initSub {
		if mqInstance.subscriber == nil {
			// TODO 應該要可以多個subID
			subscription := mqInstance.client.Subscription(t.gcpSubIDs[0])
			mqInstance.subscriber = sub.NewSubscriber(mqInstance.ctx, subscription)
		}
	}
	if t.initPub {
		if mqInstance.publisher == nil {
			mqInstance.publisher = pub.NewPublish(mqInstance.ctx, mqInstance.client.Topic(t.topicID))
		}
	}

	return mqInstance, nil
}
