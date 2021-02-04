package main

import "fmt"

// Add 宣告一種 func 型別可以傳入兩個 int 型別的值並返回一個 int 型別的值
type Add func(a int, b int) int

func process(adder Add) int {
	return adder(1, 2)
}

func main() {
	fmt.Println(process(func(a int, b int) int {
		return a + b
	}))
}
