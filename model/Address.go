package model

import (
	"errors"
	"strings"
)

var (
	ErrCityIsEmpty = errors.New("city cannot be empty")
)

type Address struct {
	CountryCode string
	PostalCode  string
	City        string
	Street      string
	HouseNumber string
}

func NewAddress(countryCode string, postalCode string, city string, street string, houseNumber string) (*Address, error) {
	address := &Address{
		CountryCode: strings.TrimSpace(countryCode),
		PostalCode:  strings.TrimSpace(postalCode),
		City:        strings.TrimSpace(city),
		Street:      strings.TrimSpace(street),
		HouseNumber: strings.TrimSpace(houseNumber),
	}

	if err := address.validate(); err != nil {
		return nil, err
	}

	return address, nil
}

func (address *Address) validate() error {
	if address.City == "" {
		return ErrCityIsEmpty
	}

	return nil
}
