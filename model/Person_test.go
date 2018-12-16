package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RegisterPerson(t *testing.T) {
	name, _ := NewName("first", "last")
	email, _ := NewEmailAddress("myemail@dot.com")

	p := Register(
		GenerateNewPersonId(),
		name,
		email,
	)

	assert.Len(t, p.RecordedEvents(), 1)
	assert.Equal(t, PersonRegisteredEventName, p.RecordedEvents()[0].eventName)
	assert.IsType(t, new(PersonRegistered), p.RecordedEvents()[0].Payload())
}

func Test_ConfirmEmailAddress(t *testing.T) {
	// given
	name, _ := NewName("first", "last")
	email, _ := NewEmailAddress("myemail@dot.com")

	p := Register(
		GenerateNewPersonId(),
		name,
		email,
	)

	// when
	p.ConfirmEmailAddress()
	p.ConfirmEmailAddress() // second call will not produce another event

	// then
	assert.Len(t, p.RecordedEvents(), 2)
}

func Test_Reconstitute(t *testing.T) {
	var events [2]*DomainEvent
	name, _ := NewName("first", "last")
	email, _ := NewEmailAddress("myemail@dot.com")

	// given
	events[0] = NewDomainEvent(
		PersonRegisteredEventName,
		NewPersonRegistered(
			NewPersonIdWithoutValidation("testingpersonid"),
			name,
			email,
		),
	)

	// when
	//p := Reconstitute(events)  TODO continue here

	// then
}
