package events

import (
	"log"
	"sync"
)

type EventBus struct {
	subscribers map[string][]chan Event
	lock        sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan Event),
	}
}

func (eb *EventBus) Subscribe(eventType string) chan Event {
	eb.lock.Lock()
	defer eb.lock.Unlock()

	ch := make(chan Event, 10) // Buffered channel to avoid blocking
	eb.subscribers[eventType] = append(eb.subscribers[eventType], ch)

	log.Printf("Subscribed to event type: %s", eventType)
	return ch
}

func (eb *EventBus) Publish(event Event) {
	eb.lock.RLock()
	defer eb.lock.RUnlock()

	subscribers, exists := eb.subscribers[string(event.Type)]
	if !exists {
		log.Printf("No subscribers for event type: %s", event.Type)
		return
	}

	for _, ch := range subscribers {
		go func(ch chan Event) {
			ch <- event
		}(ch)
	}

	log.Printf("Published event: %+v", event)
}
