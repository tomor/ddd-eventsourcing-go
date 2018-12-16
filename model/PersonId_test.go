package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateNewPersonId(t *testing.T) {
	personid := GenerateNewPersonId()

	assert.Equal(t, "TestLastname", personid.ID)
}
