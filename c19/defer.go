package main

import "fmt"

func demoDefer() {
	fmt.Println("line1")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("line2")
}

func demoNoneDefer() {
	fmt.Println("line1")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("line2")
}

func main() {
	// 沒有 defer 的情況
	fmt.Println("===== 沒 defer 的情況 =====")
	demoNoneDefer()

	// 驗證一下 defer 在何時執行
	fmt.Println("===== 以下是有 defer 的情況 =====")
	demoDefer()

	// 理論上沒有 defer 的情況下 , 會先看到 line1 然後按序印出 0~4 最後在印出 line2
	// 但其實是會印出 line1 接著在印出 line2 , 最後才打印 4~0
	// 而且是 4~0 並不是 0~4 , 這邊是有原因的 , 因為 defer 是 LIFO , 也就是說原本應該是要印 0~4 , 但 defer 的 LIFO 機制 , 讓 0 變成最先印的變成最後印

	// tips : 可以假想它是一個 LIFO 的 queue , 然後每次到 defer 要執行的時候 , 它會先把這個動作丟進 queue , 然後在最後 function 要 return 的時候 , 把 queue 的動作拿出來執行
}
