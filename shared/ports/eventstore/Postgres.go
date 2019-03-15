package eventstore

import (
	"database/sql"
	"github.com/tomor/ddd-eventsourcing-go/shared"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db}
}

func (p *Postgres) Append([]shared.DomainEvent) error {
	return nil
}

func (p *Postgres) FetchEventsByAggregateID(id shared.AggregateID) ([]shared.DomainEvent, error) {
	return nil, nil
}