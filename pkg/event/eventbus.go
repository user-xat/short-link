package event

const (
	EventLinkVisited = "link.visited"
	EventLinkUpdated = "link.updated"
	EventLinkDeleted = "link.deleted"
)

type Event struct {
	Type string
	Data any
}

type Delegate func(*Event)

type EventBus struct {
	bus         chan Event
	subscribers []Delegate
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus:         make(chan Event),
		subscribers: make([]Delegate, 0),
	}
}

func (e *EventBus) Publish(event Event) {
	e.bus <- event
}

func (e *EventBus) Subscribe(del Delegate) {
	e.subscribers = append(e.subscribers, del)
}

func (e *EventBus) Start() {
	for msg := range e.bus {
		for _, sub := range e.subscribers {
			sub(&msg)
		}
	}
}
