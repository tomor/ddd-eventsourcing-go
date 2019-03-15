package event

import (
	"github.com/tomor/ddd-eventsourcing-go/person/domain/value"
)

const (
	PersonHomeAddressAddedEventName = "PersonHomeAddressAdded"
)

type PersonHomeAddressAdded struct {
	PersonID    string `json:"person_id"`
	CountryCode string `json:"country_code"`
	PostalCode  string `json:"postal_code"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
}

func NewPersonHomeAddressAdded(personId *value.PersonID, address *value.Address) *PersonHomeAddressAdded {
	return &PersonHomeAddressAdded{
		PersonID:    personId.Value,
		CountryCode: address.CountryCode,
		PostalCode:  address.PostalCode,
		City:        address.City,
		Street:      address.Street,
		HouseNumber: address.HouseNumber,
	}
}
