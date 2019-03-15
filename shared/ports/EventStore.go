package ports

import "github.com/tomor/ddd-eventsourcing-go/shared"

type EventStore interface {
	Append([]shared.DomainEvent) error
	FetchEventsByAggregateID(id shared.AggregateID) ([]shared.DomainEvent, error)
}
