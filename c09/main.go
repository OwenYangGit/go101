package main

import "fmt"

// Person is a super
type Person struct {
	Name string
}

// Saiyan is inheritant
type Saiyan struct {
	*Person
	Power int
	Sex   string
}

func main() {
	fmt.Println("展示 struct 類似繼承的做法")

	goku := &Saiyan{
		Person: &Person{Name: "Test"},
		Power:  9000,
		Sex:    "Male",
	}

	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)

	// 都會印出 Test , 可以看到 goku 也有 Name 屬性 , 該屬性也繼承了 Person 的 Name
}
