package value

import (
	"errors"
	"github.com/satori/go.uuid"
	"strings"
)

var (
	ErrPersonIDEmptyInput = errors.New("personID cannot be empty")
)

type PersonID struct {
	Value string
}

func NewPersonId(idstring string) (*PersonID, error) {
	pid := &PersonID{strings.TrimSpace(idstring)}

	if err := pid.validate(); err != nil {
		return nil, err
	}

	return pid, nil
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

func (pid *PersonID) validate() error {
	if pid.Value == "" {
		return ErrPersonIDEmptyInput
	}

	return nil
}
