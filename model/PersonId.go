package model

import (
	"github.com/satori/go.uuid"
)

type PersonId struct {
	Value string
}

func GenerateNewPersonId() *PersonId {
	u1, err := uuid.NewV4()
	if err != nil {
		panic("generating uuid produced an error:" + err.Error())
	}

	return &PersonId{Value: u1.String()}
}

func NewPersonIdWithoutValidation(id string) *PersonId {
	return &PersonId{Value: id}
}
