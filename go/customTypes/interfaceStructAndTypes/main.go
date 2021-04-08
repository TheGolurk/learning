package main

import (
	"fmt"

	"github.com/TheGolurk/customTypes/organization"
)

func main() {
	//var p organization.Identifiable =
	/*p := organization.Person{
		FirstName: "Chris",
		LastName:  "Hdz",
	}*/
	p := organization.NewPerson("a", "b", organization.NewEuropeanUnionIdentifier("123-123-123", "Germany"))
	err := p.SetTwitterHanlder("@golurk ")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("Error %s", err.Error())
	}

	name1 := Name{First: "james", Last: "Wilson"}
	name2 := Name{First: "james", Last: "Wilson"}
	if name1 == name2 {
		println("we match")
	}

	ssn := organization.NewSocialSecurityNumber("123-12-123")
	eu := organization.NewEuropeanUnionIdentifier("12345", "Mexico")
	eu2 := organization.NewEuropeanUnionIdentifier("12345", "Mexico")

	if eu2 == eu {
		println("we match")
	}

	fmt.Printf("%T %T\n", ssn, eu)

	name3 := Name{First: "", Last: ""}
	if name3 == (Name{}) {
		println("match nil name")
	}
}

type Name struct {
	First string
	Last  string
}

type OtherName struct {
	First string
	Last  string
}
