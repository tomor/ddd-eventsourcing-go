package domain

import (
	"github.com/tomor/ddd-eventsourcing-go/person/domain/event"
	"github.com/tomor/ddd-eventsourcing-go/person/domain/value"
)

type Person struct {
	personId     *value.PersonID
	name         *value.Name
	emailAddress *value.EmailAddress
	homeAddress  *value.Address

	recordedEvents []*event.PersonDomainEvent
}

func Register(personId *value.PersonID, name *value.Name, emailAddress *value.EmailAddress) *Person {
	ev := event.NewPersonDomainEvent(
		event.PersonRegisteredEventName,
		personId,
		event.NewPersonRegistered(personId, name, emailAddress),
	)

	p := &Person{}
	p.recordThat(ev)

	return p
}

func (p *Person) ConfirmEmailAddress() {
	if !p.emailAddress.Confirmed {
		ev := event.NewPersonDomainEvent(
			event.PersonEmailAddresConfirmedEventName,
			p.personId,
			event.NewPersonEmailAddressConfirmed(p.personId),
		)

		p.recordThat(ev)
	}
}

func (p *Person) AddHomeAddress(address *value.Address) {
	ev := event.NewPersonDomainEvent(
		event.PersonHomeAddressAddedEventName,
		p.personId,
		event.NewPersonHomeAddressAdded(p.personId, address),
	)

	p.recordThat(ev)
}

func (p *Person) ChangeHomeAddress(address *value.Address) {
	ev := event.NewPersonDomainEvent(
		event.PersonHomeAddressChangedEventName,
		p.personId,
		event.NewPersonHomeAddressChanged(p.personId, address),
	)

	p.recordThat(ev)
}

/*************** Event sourcing - technical methods */

func Reconstitute(events []*event.PersonDomainEvent) *Person {
	p := &Person{}

	for _, event := range events {
		p.apply(event)
	}

	return p
}

func (p *Person) RecordedEvents() []*event.PersonDomainEvent {
	return p.recordedEvents
}

func (p *Person) recordThat(event *event.PersonDomainEvent) {
	p.recordedEvents = append(p.recordedEvents, event)
	p.apply(event) // state transition
}

func (p *Person) apply(domainEvent *event.PersonDomainEvent) {
	switch domainEvent.Meta.EventName {
	case event.PersonRegisteredEventName:
		p.whenPersonRegistered(domainEvent.Payload.(*event.PersonRegistered))
	case event.PersonEmailAddresConfirmedEventName:
		p.whenPersonEmailAddressConfirmed()
	case event.PersonHomeAddressAddedEventName:
		p.whenPersonHomeAddressAdded(domainEvent.Payload.(*event.PersonHomeAddressAdded))
	case event.PersonHomeAddressChangedEventName:
		p.whenPersonHomeAddressChanged(domainEvent.Payload.(*event.PersonHomeAddressChanged))
	default:
		// maybe error or ??..
	}
}

func (p *Person) whenPersonRegistered(domainEvent *event.PersonRegistered) {
	p.personId = value.NewPersonIdWithoutValidation(domainEvent.PersonID)
	p.name = value.NewNameWithoutValidation(domainEvent.FirstName, domainEvent.LastName)
	p.emailAddress = value.NewEmailAddressWithoutValidation(domainEvent.EmailAddress)
}

func (p *Person) whenPersonEmailAddressConfirmed() {
	p.emailAddress = p.emailAddress.Confirm()
}

func (p *Person) whenPersonHomeAddressAdded(event *event.PersonHomeAddressAdded) {
	p.homeAddress = value.NewAddressWithoutValidation(
		event.CountryCode,
		event.PostalCode,
		event.City,
		event.Street,
		event.HouseNumber,
	)
}

func (p *Person) whenPersonHomeAddressChanged(event *event.PersonHomeAddressChanged) {
	p.homeAddress = value.NewAddressWithoutValidation(
		event.CountryCode,
		event.PostalCode,
		event.City,
		event.Street,
		event.HouseNumber,
	)
}
