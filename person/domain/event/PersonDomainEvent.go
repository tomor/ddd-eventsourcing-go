package event

import (
	"encoding/json"
	"github.com/tomor/ddd-eventsourcing-go/shared"
	"time"
)

// this struct seems to be generic, maybe we could move it somewhere from person package
type PersonDomainEvent struct {
	Meta     *EventMeta         `json:"meta"`
	PersonID shared.AggregateID `json:"aggregate_id"`
	Payload  interface{}        `json:"payload"`
}

type EventMeta struct {
	EventName string    `json:"event_name"`
	OccuredAt time.Time `json:"occured_at"`
}

func NewPersonDomainEvent(eventName string, personID shared.AggregateID, payload interface{}) *PersonDomainEvent {
	return &PersonDomainEvent{
		Meta: &EventMeta{
			EventName: eventName,
			OccuredAt: time.Now(),
		},
		PersonID: personID,
		Payload:  payload,
	}
}


func (de *PersonDomainEvent) EventName() string {
	return de.Meta.EventName
}

func (de *PersonDomainEvent) OccuredAt() string {
	return de.Meta.OccuredAt.Format(time.RFC3339)
}

func (de *PersonDomainEvent) AggregateID() shared.AggregateID {
	return de.PersonID
}

func (de *PersonDomainEvent) Marshal() (string, error) {
	res, err := json.Marshal(de)

	if err != nil {
		return "", err
	}

	return string(res), nil
}
