package shared

type DomainEvent interface {
	EventName() string
	AggregateID() AggregateID
	OccuredAt() string
	Payload() interface{}
	Marshal() (string, error) // this should be done by another class ideally
}
