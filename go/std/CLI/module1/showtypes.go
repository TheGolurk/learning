package main

import "fmt"
import "reflect"

func main() {

	type person struct {
		id    int
		fname string
	}
	type employee struct {
		fname string
		age   int
	}
	nperson := person{
		id:    1,
		fname: "asd",
	}
	nempl := employee{
		fname: "aaaa",
		age:   111,
	}

	fmt.Println(nperson.id, nperson.fname)
	fmt.Println(reflect.TypeOf(nperson))
	fmt.Println(reflect.ValueOf(nperson))
	fmt.Println(reflect.ValueOf(nperson).Kind())

	addPerson(nempl)
	addPerson(nperson)

}

func addPerson(p interface{}) bool {
	if reflect.ValueOf(p).Kind() == reflect.Struct {
		v := reflect.ValueOf(p)

		switch reflect.TypeOf(p).Name() {
		case "person":
			fmt.Println("PERSON SQL QUERY")
			fmt.Println(v.Field(1))
		case "employee":
			fmt.Println("EMPLOYEE SQL QUERY")
			fmt.Println(v.Field(1))
		}

		return true
	}
	return false
}
