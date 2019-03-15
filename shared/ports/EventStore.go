package ports

type EventStore interface {
	Append([]DomainEvent) error
	FetchEventsByAggregateID(id AggregateID) ([]DomainEvent, error)
}
