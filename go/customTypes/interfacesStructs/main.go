package main

import (
	"fmt"

	"github.com/TheGolurk/pluralsight/customTypes/interfacesStructs/organization"
)

func main() {
	//var p organization.Identifiable =
	/*p := organization.Person{
		FirstName: "Chris",
		LastName:  "Hdz",
	}*/
	p := organization.NewPerson("a", "b")
	err := p.SetTwitterHanlder("@golurk ")
	if err != nil {
		fmt.Printf("Error %s", err.Error())
	}
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.ID())
	fmt.Println(p.FullName())
}
