package event

import (
	"github.com/tomor/ddd-eventsourcing-go/person/domain/value"
)

const (
	PersonEmailAddresConfirmedEventName = "PersonEmailAddressConfirmed"
)

type PersonEmailAddressConfirmed struct {
	PersonID string `json:"person_id"`
}

func NewPersonEmailAddressConfirmed(id *value.PersonID) *PersonEmailAddressConfirmed {
	return &PersonEmailAddressConfirmed{PersonID: id.Value}
}
