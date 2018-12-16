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

func NewPersonRegistered(id *PersonId, name *Name, emailAddress *EmailAddress) *PersonRegistered {
	return &PersonRegistered{
		personId:     id.ID,
		firstName:    name.FirstName,
		lastName:     name.LastName,
		emailAddress: emailAddress.Value,
	}
}

//func EventType() string {
//	return // TODO with reflection
//}
