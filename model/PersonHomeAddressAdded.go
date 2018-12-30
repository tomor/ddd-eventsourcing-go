package model

const (
	PersonHomeAddressAddedEventName = "PersonHomeAddressAdded"
)

type PersonHomeAddressAdded struct {
	personId    string
	countryCode string
	postalCode  string
	city        string
	street      string
	houseNumber string
}

func NewPersonHomeAddressAdded(personId *PersonId, address *Address) *PersonHomeAddressAdded {
	return &PersonHomeAddressAdded{
		personId:    personId.Value,
		countryCode: address.CountryCode,
		postalCode:  address.PostalCode,
		city:        address.City,
		street:      address.Street,
		houseNumber: address.HouseNumber,
	}
}
