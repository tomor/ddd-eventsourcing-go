package event

import (
	"github.com/tomor/ddd-eventsourcing-go/person/domain/value"
)

const (
	PersonRegisteredEventName = "PersonRegistered"
)

type PersonRegistered struct {
	PersonID     string `json:"person_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
}

func NewPersonRegistered(personId *value.PersonID, name *value.Name, emailAddress *value.EmailAddress) *PersonRegistered {
	return &PersonRegistered{
		PersonID:     personId.Value,
		FirstName:    name.FirstName,
		LastName:     name.LastName,
		EmailAddress: emailAddress.Value,
	}
}

//func EventType() string {
//	return // TODO with reflection
//}
