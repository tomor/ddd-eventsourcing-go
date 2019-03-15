package ports

import "github.com/tomor/ddd-eventsourcing-go/shared"

type DomainEventUnmarshaler interface {
	UnmarshalFromJSON(eventJSON string, eventName string) (shared.DomainEvent, error)
}
