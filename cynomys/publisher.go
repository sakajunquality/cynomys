package cynomys

import (
	"context"

	"google.golang.org/grpc"

	pb "google.golang.org/genproto/googleapis/pubsub/v1"
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
func (p *Publisher) Publish() error {
	req := pb.PublishRequest{
		Topic: p.Topic,
		// todo message
	}

	c, err := grpc.Dial(DialAddress, grpc.WithPerRPCCredentials(p.PerRPCCredential), grpc.WithTransportCredentials(p.TLSCredential))
	if err != nil {
		return err
	}
	defer c.Close()

	client := pb.NewPublisherClient(c)
	_, err = client.Publish(p.ctx, &req)
	if err != nil {
		return err
	}

	return nil
}

// NewMessage is ...
func NewMessage() {
}
