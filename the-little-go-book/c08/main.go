package main

import (
	"fmt"
	"reflect"
)

// Saiyan is human propeties
type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func main() {
	test := &Saiyan{
		Name:  "Test",
		Power: 9000,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  5000,
			Father: nil,
		},
	}
	fmt.Println(test.Name)
	fmt.Println(test.Power)
	fmt.Println(test.Father)
	fmt.Println(reflect.TypeOf(test.Father))
}
