package main

import (
	"fmt"
	"reflect"
)

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
	// 可以看到 Name 的值是 Goku
	fmt.Println(goku4.Name)

	// Pointer 練習

	// 先來看看沒有 Pointer 的情況
	fmt.Println(goku4.Power)
	Super1(goku4)
	fmt.Println(goku4.Power) // 看看會不會因為 Super1(goku) 調用後而改變上面已經宣告的 goku4.Power 的值

	// 來看看代入 pointer 的概念後 (&符號為 "取地址符")
	goku5 := &Saiyan{Name: "Goku"}
	goku5.Power = 9000
	fmt.Println(goku5.Power)
	Super2(goku5)
	fmt.Println(goku5.Power)

	// 來看看 pointer &符號的效果
	fmt.Println(&goku5) //看到 goku5 的 address

	// 傳遞給 Super 函數的參數仍然是一個 goku 的複本 , 只是這個複本是一個 address , 以下嘗試改變 goku6 原本的 address 來驗證傳入的 goku6 僅是一個複本
	// 透過嘗試直接改變原本的值
	goku6 := &Saiyan{Name: "Goku"}
	goku6.Power = 9000
	fmt.Println(goku.Power)
	Super3(goku6)
	fmt.Println("打印原本 goku6 的 Power")
	fmt.Println(goku6.Power) // 還是會 = 9000

	// 自己亂嘗試的
	goku7 := Saiyan{Name: "Goku7"} // 宣告一個 struct
	goku7.Power = 7000
	goku7Ptr := &goku7
	fmt.Println(reflect.TypeOf(goku7))
	fmt.Println(&goku7Ptr) // 印出 goku7 的 address
	*&goku7.Name = "Test1" // *符號 代表 pointer , 而 pointer 只能給一個 address , 而 address 只能透過 & 符號取得
	fmt.Println(goku7.Name)
}

// Super1 is a without pointer demo function
func Super1(s Saiyan) {
	s.Power = 1000
}

// Super2 demo about using pointer , and what happend
func Super2(s *Saiyan) {
	s.Power += 10000
}

// Super3 show about pointer address concept
func Super3(s *Saiyan) {
	// 將傳遞進來的 goku6 改變 address (重新指派一個新的 Saiyan)
	s = &Saiyan{Name: "Gohan", Power: 1000}
	fmt.Println("打印複本的 goku6.Power")
	fmt.Println(s.Power)
}
