package events

import (
	"time"

	"github.com/google/uuid"
)

type EventPublisher struct {
	Rate     time.Duration
	Messages chan string
}

func CreateEventPublisher(rate time.Duration) *EventPublisher {
	return &EventPublisher{Rate: rate, Messages: make(chan string)}
}

func (e *EventPublisher) Poll() <-chan string {
	go func() {
		for {
			e.Messages <- uuid.NewString()
			time.Sleep(e.Rate)
		}
	}()
	return e.Messages
}
