package main

import (
	"fmt"
	"os"
)

func main() {
	// 透過 os package 開檔 , 這個檔案叫做 test.txt
	file, err := os.Open("test.txt")
	// 如果開檔出錯 , 打印錯誤訊息 , 並執行 return -> 注意到這邊 , 這裡最後並沒有執行 file.Close() , 導致檔案其實一直是開著的
	if err != nil {
		fmt.Println(err)
		return
	}
	// 最後我們加上了 defer 呼叫 file.Close() , 因此 file.Close() 這個動作會在這個函式 return 之前執行 file.Close() 這個動作 , 至此 , 開檔的資源就被釋放掉了
	defer file.Close()
}
