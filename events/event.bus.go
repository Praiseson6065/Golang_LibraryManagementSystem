package events

import (
	"log"
	"sync"
)

type EventBus struct {
	subscribers map[string][]chan Event
	responses   map[string][]chan interface{}
	lock        sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan Event),
		responses:   make(map[string][]chan interface{}),
	}
}

func (eb *EventBus) Subscribe(eventType string) (chan Event, chan interface{}) {
	eb.lock.Lock()
	defer eb.lock.Unlock()

	eventCh := make(chan Event, 10)
	respCh := make(chan interface{}, 10)

	eb.subscribers[eventType] = append(eb.subscribers[eventType], eventCh)
	eb.responses[eventType] = append(eb.responses[eventType], respCh)

	log.Printf("Subscribed to event type: %s", eventType)
	return eventCh, respCh
}

func (eb *EventBus) Publish(event Event) []interface{} {
	eb.lock.RLock()
	defer eb.lock.RUnlock()

	subscribers, exists := eb.subscribers[string(event.Type)]
	resCh, resExists := eb.responses[string(event.Type)]
	if !exists || !resExists {
		log.Printf("No subscribers for event type: %s", event.Type)
		return nil
	}
	responseData := make([]interface{}, 0)
	var wg sync.WaitGroup
	for i, ch := range subscribers {
		wg.Add(1)
		go func(ch chan Event, respCh chan interface{}) {
			defer wg.Done()
			ch <- event
			resp := <-respCh // Wait for response
			responseData = append(responseData, resp)
		}(ch, resCh[i])
	}
	wg.Wait()

	log.Printf("Published event: %+v", event)
	return responseData
}
