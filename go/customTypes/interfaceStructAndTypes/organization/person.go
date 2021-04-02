package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandler = string

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName      string
	lastName       string
	twitterHanlder TwitterHandler
}

func NewPerson(fname, lname string) Person {
	return Person{
		firstName: fname,
		lastName:  lname,
	}
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p Person) ID() string {
	return "12345"
}

func (p *Person) SetTwitterHanlder(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHanlder = handler
	} else if !strings.HasPrefix(handler, "@") {
		return errors.New("Twitter handler must statr with an @ symbol")
	}

	p.twitterHanlder = handler
	return nil
}

func (p Person) TwitterHandler() string {
	return p.twitterHanlder
}
