package model

const (
	HomeAddressAddedEventName = "HomeAddressAdded"
)

type HomeAddressAdded struct {
	personId string
	countryCode string
	postalCode  string
	city        string
	street      string
	houseNumber string
}

func NewHomeAddressAdded(personId *PersonId, address *Address) *HomeAddressAdded  {
	return &HomeAddressAdded{
		personId:    personId.Value,
		countryCode: address.CountryCode,
		postalCode:  address.PostalCode,
		city:        address.City,
		street:      address.Street,
		houseNumber: address.HouseNumber,
	}
}