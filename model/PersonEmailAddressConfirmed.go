package model

const (
	PersonEmailAddresConfirmedEventName = "PersonEmailAddressConfirmed"
)

type PersonEmailAddressConfirmed struct {
	personId string
}

func NewPersonEmailAddressConfirmed(id *PersonId) *PersonEmailAddressConfirmed {
	return &PersonEmailAddressConfirmed{personId: id.ID}
}
