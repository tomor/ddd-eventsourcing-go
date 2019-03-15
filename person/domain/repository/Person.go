package repository

import "github.com/tomor/ddd-eventsourcing-go/person/domain"

type Person interface {
	Save(persons domain.Person) error
	FetchById(id domain.PersonID) domain.Person
}