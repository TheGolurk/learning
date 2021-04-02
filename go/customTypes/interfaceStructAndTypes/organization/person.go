package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Handler struct {
	handle string
	name   string
}

func (h Handler) randomFunc() {}

type TwitterHandler = Handler

func (th TwitterHandler) RedirectURL() string {
	return ""
}

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
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("Twitter handler must statr with an @ symbol")
	}

	p.twitterHanlder = handler
	return nil
}

func (p Person) TwitterHandler() TwitterHandler {
	return p.twitterHanlder
}
