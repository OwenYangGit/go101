package main

import "fmt"

func main() {
	fmt.Println("練習 Array")
	var scores [10]int
	scores[0] = 100
	// 上面定義有 10 個 index 的 array , 可以在 [0]~[9] 的 index 賦予值 , 因此如果在 [10] 賦值 , 則編譯會出錯

	scores2 := [4]int{9001, 9333, 212, 33}
	// 遍歷 Array
	for index, value := range scores2 {
		fmt.Println(index)
		fmt.Println(value)
	}
}
