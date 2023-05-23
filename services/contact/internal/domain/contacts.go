package domain

type Contact struct {
	Id          int
	fullName    fullName
	PhoneNumber int
}

type fullName struct {
	LastName   string
	FirstName  string
	FatherName string
}

func (c Contact) FullName() fullName {
	return c.fullName
}

func NewContact(last, first, middle string) *Contact {
	fname := fullName{
		LastName:   last,
		FirstName:  first,
		FatherName: middle,
	}

	return &Contact{fullName: fname}
}
