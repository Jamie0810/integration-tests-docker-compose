package pub

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/pkg/errors"
)

type Publisher struct {
	ctx                context.Context
	topic              *pubsub.Topic
	defaultErrorHandle func(err error, requestID string)
	errorHook          func(err error, requestID string)
}

func NewPublish(ctx context.Context, topic *pubsub.Topic) *Publisher {
	return &Publisher{
		ctx:   ctx,
		topic: topic,
		defaultErrorHandle: func(err error, requestID string) {
			log.Printf("Publish id %v error: %v", requestID, errors.WithStack(err))
		},
	}
}

func (p *Publisher) Publish(msg []byte, requestID string) error {
	result := p.topic.Publish(p.ctx, &pubsub.Message{
		Data: msg,
	})
	id, err := result.Get(p.ctx)
	log.Printf("Published a message: %v ; msg ID: %v ; requestID: %v\n", string(msg), id, requestID)
	if err != nil {
		p.triggerErrorHook(errors.WithStack(err), requestID)
		return err
	}
	return nil
}

func (p *Publisher) PublishJSON(msg interface{}, requestID string) error {
	b, err := json.Marshal(msg)
	if err != nil {
		p.triggerErrorHook(errors.WithStack(err), requestID)
		return errors.WithStack(err)
	}
	return p.Publish(b, requestID)
}

func (sb *Publisher) Options(opts ...PublisherOption) *Publisher {
	for _, opt := range opts {
		opt(sb)
	}
	return sb
}

func (sb *Publisher) triggerErrorHook(err error, requestID string) {
	if sb.errorHook != nil {
		sb.errorHook(err, requestID)
	} else {
		sb.defaultErrorHandle(err, requestID)
	}
}
