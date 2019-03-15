package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewAddress(t *testing.T) {
	address, error := NewAddress(
		"CountryCode",
		"",
		"city test name",
		"",
		"",
	)

	assert.NoError(t, error, "it should create a new address")
	assert.Equal(t, "city test name", address.City)
}

func Test_NewAddress_EmptyCity(t *testing.T) {
	_, error := NewAddress(
		"",
		"",
		"",
		"",
		"",
	)

	assert.EqualError(t, error, ErrCityIsEmpty.Error())
}
