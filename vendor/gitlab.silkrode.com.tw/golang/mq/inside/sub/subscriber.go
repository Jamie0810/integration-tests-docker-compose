package sub

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

const (
	syncMode                      = "sync"
	asyncMode                     = "async"
	defaultMaxOutstandingMessages = 20
)

type Subscriber struct {
	ctx                context.Context
	subscription       *pubsub.Subscription
	pullMode           string
	defaultErrorHandle func(err error, msg *pubsub.Message)
	errorHook          func(err error, msgData string, msgID string)
}

func NewSubscriber(ctx context.Context, subscription *pubsub.Subscription) *Subscriber {
	subscription.ReceiveSettings.MaxOutstandingMessages = defaultMaxOutstandingMessages
	return &Subscriber{
		ctx:          ctx,
		subscription: subscription,
		defaultErrorHandle: func(err error, msg *pubsub.Message) {
			errorLog := ""
			if msg != nil {
				errorLog = fmt.Sprintf("mq err: %s,msgData: %s,msgID: %s", err.Error(), msg.Data, msg.ID)
			} else {
				errorLog = fmt.Sprintf("mq err: %s", err.Error())
			}
			log.Print(errorLog)
		},
	}
}

func (sb *Subscriber) Options(opts ...SubscriberOption) *Subscriber {
	for _, opt := range opts {
		opt(sb)
	}
	return sb
}

func (sb *Subscriber) Subscribe(process func(context.Context, []byte, string) error) {
	var err error
	if sb.pullMode == syncMode {
		sb.subscription.ReceiveSettings.Synchronous = true
		err = sb.pullMessageSync(sb.ctx, process)
	}
	if sb.pullMode == asyncMode {
		sb.subscription.ReceiveSettings.Synchronous = false
		err = sb.pullMessageAsync(sb.ctx, process)
	}
	if err != nil {
		sb.triggerErrorHook(err, nil)
	}
}

func (sb *Subscriber) pullMessageSync(ctx context.Context, process func(context.Context, []byte, string) error) error {
	cm := make(chan *pubsub.Message)
	go func() {
		for {
			select {
			case msg := <-cm:
				err := process(sb.ctx, msg.Data, msg.ID)
				if err != nil {
					sb.triggerErrorHook(err, msg)
				}
				msg.Ack()
			case <-ctx.Done():
				return
			}
		}
	}()
	return sb.subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		cm <- msg
	})
}

func (sb *Subscriber) pullMessageAsync(ctx context.Context, process func(context.Context, []byte, string) error) error {
	return sb.subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if err := process(sb.ctx, msg.Data, msg.ID); err != nil {
			sb.triggerErrorHook(err, msg)
		}
		msg.Ack()
	})
}

func (sb *Subscriber) triggerErrorHook(err error, msg *pubsub.Message) {
	if sb.errorHook != nil {
		if msg == nil {
			sb.errorHook(err, "", "")
		} else {
			sb.errorHook(err, string(msg.Data), msg.ID)
		}
	} else {
		sb.defaultErrorHandle(err, msg)
	}
}
