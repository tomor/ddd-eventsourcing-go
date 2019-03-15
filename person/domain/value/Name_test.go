package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewName(t *testing.T) {
	name, error := NewName("TestFirstname", "TestLastname")

	assert.NoError(t, error, "it should create a new Name")
	assert.Equal(t, "TestFirstname", name.FirstName)
	assert.Equal(t, "TestLastname", name.LastName)
}

func Test_NewName_InvalidFirstName(t *testing.T) {
	_, error := NewName("", "TestLastname")

	assert.EqualError(t, error, ErrFistNameIsEmpty.Error())
}

func Test_NewName_InvalidLastName(t *testing.T) {
	_, error := NewName("First name", "")

	assert.EqualError(t, error, ErrLastNameIsEmpty.Error())
}
