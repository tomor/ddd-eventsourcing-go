package model

type Person struct {
	personId     *PersonId
	name         *Name
	emailAddress *EmailAddress
	homeAddress  *Address

	recordedEvents []*DomainEvent
}

func Register(personId *PersonId, name *Name, emailAddress *EmailAddress) *Person {
	ev := NewDomainEvent(
		PersonRegisteredEventName,
		NewPersonRegistered(personId, name, emailAddress),
	)

	p := &Person{}
	p.recordThat(ev)

	return p
}

func (p *Person) ConfirmEmailAddress() {
	if !p.emailAddress.Confirmed {
		ev := NewDomainEvent(
			PersonEmailAddresConfirmedEventName,
			NewPersonEmailAddressConfirmed(p.personId),
		)

		p.recordThat(ev)
	}
}

func (p *Person) AddHomeAddress(address *Address) {
	ev := NewDomainEvent(
		PersonHomeAddressAddedEventName,
		NewPersonHomeAddressAdded(p.personId, address),
	)

	p.recordThat(ev)
}

func (p *Person) ChangeHomeAddress(address *Address) {
	ev := NewDomainEvent(
		PersonHomeAddressChangedEventName,
		NewPersonHomeAddressChanged(p.personId, address),
	)

	p.recordThat(ev)
}

/*************** Event sourcing - technical methods */

func Reconstitute(events []*DomainEvent) *Person {
	p := &Person{}

	for _, event := range events {
		p.apply(event)
	}

	return p
}

func (p *Person) RecordedEvents() []*DomainEvent {
	return p.recordedEvents
}

func (p *Person) recordThat(event *DomainEvent) {
	p.recordedEvents = append(p.recordedEvents, event)
	p.apply(event) // state transition
}

func (p *Person) apply(event *DomainEvent) {
	switch event.EventName() {
	case PersonRegisteredEventName:
		p.whenPersonRegistered(event.Payload().(*PersonRegistered))
	case PersonEmailAddresConfirmedEventName:
		p.whenPersonEmailAddressConfirmed()
	case PersonHomeAddressAddedEventName:
		p.whenPersonHomeAddressAdded(event.Payload().(*PersonHomeAddressAdded))
	case PersonHomeAddressChangedEventName:
		p.whenPersonHomeAddressChanged(event.Payload().(*PersonHomeAddressChanged))
	default:
		// maybe error or ??..
	}
}

func (p *Person) whenPersonRegistered(event *PersonRegistered) {
	p.personId = NewPersonIdWithoutValidation(event.personId)
	p.name = NewNameWithoutValidation(event.firstName, event.lastName)
	p.emailAddress = NewEmailAddressWithoutValidation(event.emailAddress)
}

func (p *Person) whenPersonEmailAddressConfirmed() {
	p.emailAddress = p.emailAddress.Confirm()
}

func (p *Person) whenPersonHomeAddressAdded(event *PersonHomeAddressAdded) {
	p.homeAddress = NewAddressWithoutValidation(
		event.countryCode,
		event.postalCode,
		event.city,
		event.street,
		event.houseNumber,
	)
}

func (p *Person) whenPersonHomeAddressChanged(event *PersonHomeAddressChanged) {
	p.homeAddress = NewAddressWithoutValidation(
		event.countryCode,
		event.postalCode,
		event.city,
		event.street,
		event.houseNumber,
	)
}
