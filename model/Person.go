package model

type Person struct {
	id                *PersonId
	name              *Name
	basicEmailAddress *EmailAddress
	homeAddress       *Address

	recordedEvents []*DomainEvent
}

func Register(id *PersonId, name *Name, emailAddress *EmailAddress) *Person {
	ev := NewDomainEvent(
		PersonRegisteredEventName,
		NewPersonRegistered(id, name, emailAddress),
	)

	p := &Person{}
	p.recordThat(ev)

	return p
}

func (p *Person) ConfirmEmailAddress() {
	if !p.basicEmailAddress.Confirmed {
		ev := NewDomainEvent(
			PersonEmailAddresConfirmedEventName,
			NewPersonEmailAddressConfirmed(p.id),
		)

		p.recordThat(ev)
	}
}

//TODO: finish adress: events:
//AddressAdded
//AddressChanged

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
	default:
		// maybe error or ??..

	}
}

func (p *Person) whenPersonRegistered(event *PersonRegistered) {
	p.id = NewPersonIdWithoutValidation(event.personId)
	p.name = NewNameWithoutValidation(event.firstName, event.lastName)
	p.basicEmailAddress = NewEmailAddressWithoutValidation(event.emailAddress)
}

func (p *Person) whenPersonEmailAddressConfirmed() {
	p.basicEmailAddress = p.basicEmailAddress.Confirm()
}
