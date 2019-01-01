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

func Test_Reconstitute_With_HomeAddress_EmailNotConfirmed(t *testing.T) {
	// given
	personId := NewPersonIdWithoutValidation("testingpersonid")
	name := NewNameWithoutValidation("first", "last")
	email := NewEmailAddressWithoutValidation("myemail@dot.com")

	events := []*DomainEvent{
		NewDomainEvent(
			PersonRegisteredEventName,
			NewPersonRegistered(
				personId,
				name,
				email,
			),
		),
		NewDomainEvent(
			PersonEmailAddresConfirmedEventName,
			NewPersonEmailAddressConfirmed(
				personId,
			),
		),
		NewDomainEvent(
			PersonHomeAddressAddedEventName,
			NewPersonHomeAddressAdded(
				personId,
				NewAddressWithoutValidation(
					"DE",
					"80686",
					"München",
					"Our test street",
					"250",
				),
			),
		),
	}

	// when
	p := Reconstitute(events)

	// then
	assert.Equal(t, personId, p.personId)
	assert.Equal(t, email.Value, p.emailAddress.Value)
	assert.Equal(t, true, p.emailAddress.Confirmed)
	assert.Equal(t, name.FirstName, p.name.FirstName)
	assert.Equal(t, name.LastName, p.name.LastName)
	assert.Equal(t, "München", p.homeAddress.City)
	assert.Equal(t, "DE", p.homeAddress.CountryCode)
	assert.Equal(t, "250", p.homeAddress.HouseNumber)
	assert.Equal(t, "Our test street", p.homeAddress.Street)
	assert.Equal(t, "80686", p.homeAddress.PostalCode)
}

func Test_Reconstitute_With_ConfirmedEmailAddress(t *testing.T) {
	personId := NewPersonIdWithoutValidation("testingpersonid")
	name := NewNameWithoutValidation("first", "last")
	email := NewEmailAddressWithoutValidation("myemail@dot.com")

	// given
	events := []*DomainEvent{
		NewDomainEvent(
			PersonRegisteredEventName,
			NewPersonRegistered(
				personId,
				name,
				email,
			),
		),
		NewDomainEvent(
			PersonEmailAddresConfirmedEventName,
			NewPersonEmailAddressConfirmed(
				personId,
			),
		),
	}

	// when
	p := Reconstitute(events)

	// then
	assert.Equal(t, personId, p.personId)
	assert.Equal(t, email.Value, p.emailAddress.Value)
	assert.Equal(t, true, p.emailAddress.Confirmed)
	assert.Equal(t, name.FirstName, p.name.FirstName)
	assert.Equal(t, name.LastName, p.name.LastName)
	assert.Empty(t, p.homeAddress)
}

/*** test helper methods ***/
func registerPerson() *Person {
	name, _ := NewName("first", "last")
	email, _ := NewEmailAddress("myemail@dot.com")

	return Register(
		GenerateNewPersonId(),
		name,
		email,
	)
}
