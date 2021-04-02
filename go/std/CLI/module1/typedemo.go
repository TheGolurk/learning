package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	id   int
	name string
}

func main() {

	employees := make([]employee, 3)

	etype := reflect.TypeOf(employees)
	newEmployeeList := reflect.MakeSlice(etype, 0, 0)
	newEmployeeList = reflect.Append(newEmployeeList, reflect.ValueOf(employee{3, "das"}))

	fmt.Println(employees)
	fmt.Println(newEmployeeList)
}
