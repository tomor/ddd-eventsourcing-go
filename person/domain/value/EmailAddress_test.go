package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewEmailAddress(t *testing.T) {
	email, error := NewEmailAddress("testemail@dot.com")

	assert.NoError(t, error, "it should create a new email address")
	assert.Equal(t, "testemail@dot.com", email.Value)
}

func Test_NewEmailAddress_EmptyValue(t *testing.T) {
	_, error := NewEmailAddress("")

	assert.EqualError(t, error, ErrEmailAddressIsEmpty.Error())
}
