package eventstore

import (
	"database/sql"
	"github.com/tomor/ddd-eventsourcing-go/shared/ports"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db}
}

func (p *Postgres) Append([]ports.DomainEvent) error {
	return nil
}

func (p *Postgres) FetchEventsByAggregateID(id ports.AggregateID) ([]ports.DomainEvent, error) {
	return nil, nil
}