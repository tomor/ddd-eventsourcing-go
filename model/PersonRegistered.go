package model

const (
	PersonRegisteredEventName = "PersonRegistered"
)

type PersonRegistered struct {
	personId     string
	firstName    string
	lastName     string
	emailAddress string
}

func NewPersonRegistered(personId *PersonId, name *Name, emailAddress *EmailAddress) *PersonRegistered {
	return &PersonRegistered{
		personId:     personId.Value,
		firstName:    name.FirstName,
		lastName:     name.LastName,
		emailAddress: emailAddress.Value,
	}
}

//func EventType() string {
//	return // TODO with reflection
//}
