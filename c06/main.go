package main

import "fmt"

// Saiyan man data structure definition
type Saiyan struct {
	Name  string
	Power int
}

func main() {

	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	// 沒用會報錯 , 用一下
	fmt.Println(goku.Name)

	// 相反也可以先給予一個空
	goku2 := Saiyan{}
	// 沒用會報錯 , 用一下
	fmt.Println(goku2.Name)

	// 又或是
	goku3 := Saiyan{Name: "Goku"}
	goku3.Power = 9000
	fmt.Println(goku3.Power)

	// 雖然這樣用比較簡潔 , 但是不建議 , 畢竟根本看不出來第一個 Goku 的變量名稱是甚麼
	goku4 := Saiyan{"Goku", 9000}
	fmt.Println(goku4.Name)
}
