package eventstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/tomor/ddd-eventsourcing-go/person/domain/event"
	"github.com/tomor/ddd-eventsourcing-go/person/domain/value"
	"github.com/tomor/ddd-eventsourcing-go/shared/ports"
	"testing"
)

func Test_Append(t *testing.T) {
	// given
	db := initDB(t)
	es := NewPostgres(db)

	personID := value.NewPersonIdWithoutValidation("123")

	// when

	events := []ports.DomainEvent{
		event.NewPersonDomainEvent(event.PersonEmailAddresConfirmedEventName, personID, event.NewPersonEmailAddressConfirmed(personID)),
	}


	res := es.Append(events)

	// then
	assert.Equal(t, "1", res)
}

func Test_FetchEventsByAggregateID(t *testing.T) {
	// given

	// when

	// then
}

/* helper methods */

func initDB(t *testing.T) *sql.DB {
	dbURI := "postgres://admin:admin@localhost:5432/eventstore"
	db, err := sql.Open("postgres", dbURI)

	assert.NoError(t, err, "it should connect to postgres")
	if err != nil {
		panic("")
	}

	return db
}