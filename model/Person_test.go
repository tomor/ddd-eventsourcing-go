package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RegisterPerson(t *testing.T) {
	// given
	// when
	p := registerPerson()

	// then
	assert.Len(t, p.RecordedEvents(), 1)
	assert.Equal(t, PersonRegisteredEventName, p.RecordedEvents()[0].eventName)
	assert.IsType(t, new(PersonRegistered), p.RecordedEvents()[0].Payload())
}

func Test_ConfirmEmailAddress(t *testing.T) {
	// given
	p := registerPerson()

	// when
	p.ConfirmEmailAddress()
	p.ConfirmEmailAddress() // second call will not produce another event

	// then
	assert.Len(t, p.RecordedEvents(), 2)
	assert.Equal(t, PersonRegisteredEventName, p.RecordedEvents()[0].eventName)
	assert.Equal(t, PersonEmailAddresConfirmedEventName, p.RecordedEvents()[1].eventName)
	assert.IsType(t, new(PersonEmailAddressConfirmed), p.RecordedEvents()[1].Payload())
	// (Note #1) here we could also test that event contains correct values
	// it depends if we have other tests that will ensure that or not
}

func Test_AddHomeAddress(t *testing.T) {
	// given
	p := registerPerson()

	// when
	homeAddress := NewAddressWithoutValidation(
		"country code",
		"postal code",
		"city",
		"street",
		"15",
	)

	p.AddHomeAddress(homeAddress)

	// then
	assert.Len(t, p.RecordedEvents(), 2)
	assert.Equal(t, PersonRegisteredEventName, p.RecordedEvents()[0].eventName)
	assert.Equal(t, PersonHomeAddressAddedEventName, p.RecordedEvents()[1].eventName)
	assert.IsType(t, new(PersonHomeAddressAdded), p.RecordedEvents()[1].Payload())
}

// TODO tests for apply() method are still missing

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

/*** helper methods ***/
func registerPerson() *Person {
	name, _ := NewName("first", "last")
	email, _ := NewEmailAddress("myemail@dot.com")

	return Register(
		GenerateNewPersonId(),
		name,
		email,
	)
}
