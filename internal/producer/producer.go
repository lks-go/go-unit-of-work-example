package producer

import (
	"fmt"
)

func New() *QueueProducer {
	return &QueueProducer{}
}

type QueueProducer struct{}

func (p *QueueProducer) Send(email, code string) error {
	fmt.Printf("Sending message with code '%s' for email '%s' to queue", code, email)
	return nil
}
