package value

import (
	"errors"
	"strings"
)

var (
	ErrEmailAddressIsEmpty = errors.New("email address cannot be empty")
)

type EmailAddress struct {
	Value     string
	Confirmed bool
}

func NewEmailAddress(value string) (*EmailAddress, error) {
	emailadr := &EmailAddress{
		Value:     strings.TrimSpace(value),
		Confirmed: false,
	}

	if err := emailadr.validate(); err != nil {
		return nil, err
	}

	return emailadr, nil
}

func NewEmailAddressWithoutValidation(value string) *EmailAddress {
	return &EmailAddress{
		Value: strings.TrimSpace(value),
	}
}

func (em EmailAddress) Confirm() *EmailAddress {
	return &EmailAddress{Value: em.Value, Confirmed: true}
}

func (emailaddress *EmailAddress) validate() error {
	if emailaddress.Value == "" {
		return ErrEmailAddressIsEmpty
	}

	return nil
}
