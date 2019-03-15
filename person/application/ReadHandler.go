package application

import (
	"github.com/tomor/ddd-eventsourcing-go/person/domain/repository"
	"github.com/tomor/ddd-eventsourcing-go/person/domain/value"
)

func GetPersonById(idstring string) *Person {
	id, _ := value.NewPersonID(idstring)

	repository := .....
	events := repository.FetchEventsForAggregate(id)

	person :=  Reconstitute(events)

	return Person
}