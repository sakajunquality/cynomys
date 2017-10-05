package cynomys

import (
	"context"
)

const (
	// DialAddress is ...
	DialAddress = "pubsub.googleapis.com:443"
)

// Publisher is ...
type Publisher struct {
	ctx context.Context
	PerRPCCredential
	TLSCredential
	Topic    string
	Messages []Message
}

// Message is ...
type Message struct {
	ID         string
	Data       []byte
	Attributes []Attribute
}

// Attribute is ...
type Attribute struct {
	Key   string
	Value string
}

// NewPublisher is ...
func NewPublisher(ctx context.Context, topic string) Publisher {
	return Publisher{
		ctx:   ctx,
		Topic: topic,
	}
}

// Compose is ...
func (p *Publisher) Compose() {

}

// Publish is ...
func (p *Publisher) Publish() {

}

// NewMessage is ...
func NewMessage() {
}
