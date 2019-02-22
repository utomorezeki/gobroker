package gobroker

// Message encapsulates actual message being sent & published by message broker
type Message struct {
	Body     []byte
	Attempts int
}

// Handler defines how client should handle incoming messages as subscriber
type Handler func(msg *Message) error

// Implementation defines supported adapters
type Implementation int

const (
	RabbitMQ = Implementation(iota)
	Google
)
