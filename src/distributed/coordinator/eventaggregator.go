package coordinator

import "time"

// EventAggregator event aggregator
type EventAggregator struct {
	listeners map[string][]func(EventData)
}

// NewEventAggregator new event aggregator
func NewEventAggregator() *EventAggregator {
	ea := EventAggregator{
		listeners: make(map[string][]func(EventData)),
	}
	return &ea
}

// AddListener add listener
func (ea *EventAggregator) AddListener(name string, f func(EventData)) {
	ea.listeners[name] = append(ea.listeners[name], f)
}

// PublishEvent publish event
func (ea *EventAggregator) PublishEvent(name string, eventData EventData) {
	if ea.listeners[name] != nil {
		for _, r := range ea.listeners[name] {
			r(eventData)
		}
	}
}

// EventData event data
type EventData struct {
	Name      string
	Value     float64
	Timestamp time.Time
}
