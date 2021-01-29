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

	goku2 := &Saiyan{
		Person: &Person{Name: "Test2"},
		Power:  1999,
		Sex:    "Female",
	}

	goku2.Name = "Demo1"
	// goku2.Person.Name = "Demo2" // 這行會蓋掉上面那行的值

	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)
	// 都會印出 Test , 可以看到 goku 也有 Name 屬性 , 該屬性也繼承了 Person 的 Name

	fmt.Println(goku2.Name)
	fmt.Println(goku2.Person.Name)
	// 可以看到改變 goku2.Name 的值 , 也會順便改變 goku2.Person.Name 的值 , 因此這邊要注意的是 , 繼承是一種概念 , 在實作上要注意

}
