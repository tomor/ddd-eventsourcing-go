package value

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateNewPersonId(t *testing.T) {
	personid := GenerateNewPersonId()

	assert.NotEmpty(t, personid.Value)
	assert.Len(t, personid.Value, 36)
	// example: f6e57303-f68c-464d-a8c8-71336b5c0d4b
	assert.Regexp(t, regexp.MustCompile(`^([a-z0-9]+-)*[a-z0-9]+$`), personid.Value)
}
