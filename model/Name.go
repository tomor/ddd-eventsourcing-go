package model

import (
	"errors"
	"strings"
)

var (
	ErrFistNameIsEmpty = errors.New("first name cannot be empty")
	ErrLastNameIsEmpty = errors.New("last name cannot be empty")
)

type Name struct {
	FirstName string
	LastName  string
}

func NewName(firstName string, lastName string) (*Name, error) {
	name := &Name{
		FirstName: strings.TrimSpace(firstName),
		LastName:  strings.TrimSpace(lastName),
	}

	if err := name.validate(); err != nil {
		return nil, err
	}

	return name, nil
}

func NewNameWithoutValidation(firstName string, lastName string) *Name {
	return &Name{
		FirstName: strings.TrimSpace(firstName),
		LastName:  strings.TrimSpace(lastName),
	}

}

func (name *Name) validate() error {
	if name.FirstName == "" {
		return ErrFistNameIsEmpty
	}

	if name.LastName == "" {
		return ErrLastNameIsEmpty
	}

	return nil
}
