package pubsub

import (
	"bytes"
	"context"

	"gitlab.rozetka.company/goods-service/log"
)

type (
	Subscriber interface {
		Notify(ctx context.Context, body interface{}, error chan error)
		Name() string
		Close() error
	}

	Publisher interface {
		Publish(ctx context.Context, body interface{}, err chan error)
		AddSubscriber(s Subscriber)
		Stop(ctx context.Context) error
	}

	inBody struct {
		ctx       context.Context
		body      interface{}
		returnErr chan error
	}
)

type publisher struct {
	subscribers       []Subscriber
	in                chan inBody
	addSubscriberChan chan Subscriber
	stop              chan chan struct{}
}

func NewPublisher(ctx context.Context) *publisher {
	p := &publisher{
		in:                make(chan inBody),
		addSubscriberChan: make(chan Subscriber),
		stop:              make(chan chan struct{}),
	}

	go p.start(ctx)

	return p
}

func (p *publisher) Publish(ctx context.Context, body interface{}, err chan error) {
	in := inBody{
		ctx:       ctx,
		body:      body,
		returnErr: err,
	}

	p.in <- in
}

func (p *publisher) start(ctx context.Context) {
	defer log.Info(ctx, "[publisher] finish listening for messages")
	log.Info(ctx, "[publisher] start listening for messages")

	for {
		select {
		case body := <-p.in:
			for _, s := range p.subscribers {
				s.Notify(body.ctx, body.body, body.returnErr)
			}
		case s := <-p.addSubscriberChan:
			p.subscribers = append(p.subscribers, s)
		case ack := <-p.stop:
			ack <- struct{}{}
			return
		}
	}
}

func (p *publisher) AddSubscriber(s Subscriber) {
	p.addSubscriberChan <- s
}

func (p *publisher) Stop(ctx context.Context) error {
	log.Info(ctx, "publisher stopping")
	stopAck := make(chan struct{})
	p.stop <- stopAck
	<-stopAck
	log.Info(ctx, "publisher stopped")

	return nil
}

type subscribersError []error

func (s subscribersError) Error() string {
	buffer := bytes.Buffer{}

	for i, err := range s {
		if i > 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(err.Error())
	}

	return buffer.String()
}
