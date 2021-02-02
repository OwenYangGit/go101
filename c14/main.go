package main

import "fmt"

func main() {
	type Saiyan struct {
		Name  string
		Power int
	}

	a := make([]Saiyan, 10) //  len = 10 , cap = 10
	b := make([]*Saiyan, 10)
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(cap(a))
	fmt.Println(cap(b))
}
