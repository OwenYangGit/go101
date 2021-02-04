package main

import (
	"fmt"
)

// Saiyan is human propeties
type Saiyan struct {
	Name  string
	Power int
}

func (s *Saiyan) showPower() {
	fmt.Println(s.Power)
}

func (s *Saiyan) setName(name string) {
	s.Name = name
}

func main() {
	man1 := &Saiyan{Name: "Goku1", Power: 5000}
	fmt.Println(man1.Power)
	man1.Power += 1000
	man1.showPower()
	man1.setName("Op1")
	fmt.Println(man1.Name)
}
