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
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectURL())
	fmt.Println(p.ID())
	fmt.Println(p.Country())
	fmt.Println(p.FullName())
}
