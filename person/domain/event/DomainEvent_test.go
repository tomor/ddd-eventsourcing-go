package event

import (
	"github.com/stretchr/testify/assert"
	"github.com/tomor/ddd-eventsourcing-go/person/domain/value"
	"testing"
)

func Test_DomainEvent_Marshal_success(t *testing.T) {
	// given
	event := NewPersonDomainEvent(
		PersonEmailAddresConfirmedEventName,
		NewPersonEmailAddressConfirmed(value.NewPersonIdWithoutValidation("123")),
	)

	// when
	res, err := event.Marshal()

	// then
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	//assert.Equal(t, "{\"meta\":{\"event_name\":\"PersonEmailAddressConfirmed\",\"occured_at\":\"2019-03-15T14:07:38.317709247+01:00\"},\"payload\":{\"person_id\":\"123\"}}", res)
}