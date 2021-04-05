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

type TwitterHandler string

func (th TwitterHandler) RedirectURL() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twtter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type Name struct {
	first string
	last  string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHanlder TwitterHandler
}

func NewPerson(fname, lname string) Person {
	return Person{
		Name: Name{
			first: fname,
			last:  lname,
		},
	}
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
