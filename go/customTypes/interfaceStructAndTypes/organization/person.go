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

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

type erupeanUnionIdentifier struct {
	id      string
	country string
}

func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return erupeanUnionIdentifier{
		id:      id,
		country: country,
	}
}

func (eui erupeanUnionIdentifier) ID() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

func (eui erupeanUnionIdentifier) Country() string {
	return eui.country
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
	Citizen
}

func NewPerson(fname, lname string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: fname,
			last:  lname,
		},
		Citizen: citizen,
	}
}

func (p *Person) ID() string {
	return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
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
