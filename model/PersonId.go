package model

import (
	"errors"

	"github.com/satori/go.uuid"
)

var (
	ErrInvalidPersonId = errors.New("first name cannot be empty")
)

type PersonId struct {
	ID string
}

func GenerateNewPersonId() *PersonId {
	u1, err := uuid.NewV4()
	if err != nil {
		panic("generating uuid produced an error:" + err.Error())
	}

	return &PersonId{ID: u1.String()}
}

func NewPersonIdWithoutValidation(id string) *PersonId {
	return &PersonId{ID: id}
}
