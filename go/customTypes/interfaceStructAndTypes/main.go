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
	//if name1 == name2 {
	//		println("we match")
	//	}
	fmt.Println(name1, name2)

	ssn := organization.NewSocialSecurityNumber("123-12-123")
	eu := organization.NewEuropeanUnionIdentifier("12345", "Mexico")
	eu2 := organization.NewEuropeanUnionIdentifier("12345", "Mexico")

	fmt.Printf("%T %T %T\n", ssn, eu, eu2)

	//name3 := Name{First: "", Last: ""}
	//if name3 == (Name{}) {
	//		println("match nil name")
	//	}

	//portfolio := map[Name][]organization.Person{}
	//portfolio[name1] = []organization.Person{p}

	if name1.Equals(Name{}) {
		println("match equals")
	}

	fmt.Println(p.ID())
	fmt.Println(p.Country())
}

type Name struct {
	First  string
	Last   string
	Middle []string
}

func (n Name) Equals(otherName Name) bool {
	return n.First == otherName.First && n.Last == otherName.Last && len(n.Middle) == len(otherName.Middle)
}

type OtherName struct {
	First string
	Last  string
}
