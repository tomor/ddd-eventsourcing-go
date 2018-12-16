package model

type DomainEvent struct {
	eventName string
	//occuredAt
	payload interface{}
}

func NewDomainEvent(eventName string, payload interface{}) *DomainEvent {
	return &DomainEvent{
		eventName: eventName,
		payload:   payload,
	}
}

func (de *DomainEvent) EventName() string {
	return de.eventName
}

func (de *DomainEvent) Payload() interface{} {
	return de.payload
}
