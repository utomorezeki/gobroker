package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"google.golang.org/api/option"
	"log"
	"sync"
)

type googlePub struct {
	c      *pubsub.Client
	topics *sync.Map
	m      sync.Mutex
}

func newGooglePub(cfg *config) *googlePub {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, cfg.projectID, option.WithCredentialsFile(cfg.googleJSONFile))
	if nil != err {
		log.Fatalln("failed to initialize google publisher", err)
	}
	return &googlePub{
		c:      client,
		topics: &sync.Map{},
	}
}

func (g *googlePub) Publish(topic string, message interface{}) error {
	data, err := json.Marshal(message)
	if nil != err {
		return err
	}

	return g.publish(topic, &pubsub.Message{
		Data: data,
	})
}

func (g *googlePub) publish(topic string, message *pubsub.Message) error {
	ctx := context.Background()
	t := g.getTopic(topic)
	result := t.Publish(ctx, message)
	_, err := result.Get(ctx)

	return err
}

func (g *googlePub) getTopic(topic string) *pubsub.Topic {
	if t, exist := g.topics.Load(topic); exist {
		return t.(*pubsub.Topic)
	}
	g.m.Lock()
	defer g.m.Unlock()

	if t, exist := g.topics.Load(topic); exist {
		return t.(*pubsub.Topic)
	}

	t := g.c.Topic(topic)
	g.topics.Store(topic, t)

	return t
}
