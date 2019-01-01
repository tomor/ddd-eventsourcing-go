package model

const (
	PersonHomeAddressChangedEventName = "PersonHomeAddressChanged"
)

type PersonHomeAddressChanged struct {
	personId    string
	countryCode string
	postalCode  string
	city        string
	street      string
	houseNumber string
}

func NewPersonHomeAddressChanged(personId *PersonId, address *Address) *PersonHomeAddressChanged {
	return &PersonHomeAddressChanged{
		personId:    personId.Value,
		countryCode: address.CountryCode,
		postalCode:  address.PostalCode,
		city:        address.City,
		street:      address.Street,
		houseNumber: address.HouseNumber,
	}
}
