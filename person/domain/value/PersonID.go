package value

import (
	"github.com/satori/go.uuid"
)

type PersonID struct {
	Value string
}

func GenerateNewPersonId() *PersonID {
	u1 := uuid.NewV4()

	return &PersonID{Value: u1.String()}
}

func NewPersonIdWithoutValidation(id string) *PersonID {
	return &PersonID{Value: id}
}

func (pid *PersonID) ID() string {
	return pid.Value
}
