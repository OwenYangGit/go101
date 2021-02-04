package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
	fmt.Println("Arg 0 is", os.Args[0])
}

// 使用 go run main.go 9000
// os.Args[0] 在 go 預設情況下 , 這個第 0 個 index 值為 "執行檔案的路徑位置"
