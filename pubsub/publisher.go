package pubsub

import "github.com/febytanzil/gobroker"

// Publisher provides adapter to publish message
type Publisher interface {
	Publish(topic string, message interface{}) error
}

// NewPublisher implements adapter instance for Publisher
func NewPublisher(impl gobroker.Implementation, options ...Option) Publisher {
	c := &config{}
	for _, o := range options {
		o(c)
	}

	switch impl {
	case gobroker.RabbitMQ:
		return newRabbitMQPub(c)
	case gobroker.Google:
		return newGooglePub(c)
	default:
		return nil
	}
}
