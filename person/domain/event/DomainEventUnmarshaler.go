package event

import (
	"github.com/tomor/ddd-eventsourcing-go/shared"
)

type  DomainEventUnmarshaler struct {}

func (u *DomainEventUnmarshaler) UnmarshalFromJSON(eventJSON string, eventName string) (shared.DomainEvent, error) {
	switch eventName {
	case PersonRegisteredEventName:

	}
}
