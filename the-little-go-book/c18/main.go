package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	// 如果傳遞參數不是剛好為 2 個 , 以 os 的 exit code 1 跳出
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	// 接收 strconv.Atoi 函數傳回來的值 , 可以去查看 strconv.Atoi 的 function 返回值 , func Atoi(s string) (int, error)
	// 可以看到它會返回兩種 type 的值 , 第一個為 int , 第二個為 error . 所以這邊為 n 變數接收 int 型別的返回值 , err 變數接收 error 型別的返回值
	n, err := strconv.Atoi(os.Args[1])

	// 判斷 err 變數是不是為 nil  , 是的話代表有正常轉換 , 所以可以打印出 int , 否則就做錯誤處理 print 出 not a valid number
	if err != nil {
		fmt.Println("not a valid number")
		fmt.Println(reflect.TypeOf(n))
		// 以下打印兩個變數 , 可以透過故意錯誤執行 , 查看兩個變數的值的變化 , 例如若想看 err != nil 的情況 , n 跟 err 值分別為何 可以輸入 `go run main.go NotAInt`
		fmt.Println(n)
		fmt.Println(err)
	} else {
		// 同理 , 想觀察 err = nil 的情況下 , n 跟 err 變數的值為何 , 可以輸入 `go run main.go 3`
		fmt.Println(n)
		fmt.Println(err)
		fmt.Println(reflect.TypeOf(n))
	}
}
