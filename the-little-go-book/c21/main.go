package main

import "fmt"

func main() {
	stra := "the spice must flow"
	byts := []byte(stra)
	fmt.Println(byts)
	strb := string(byts)
	fmt.Println(strb)

	// 當 string 中含有 unicode 的 runes 特殊字元的時候 , 在使用 len() 計算長度時 , 返回的答案可能並不正確
	fmt.Println(len("�"))

	// 正常 string
	fmt.Println(len("string"))

	// 書中有提到這類特殊的用例 , 這邊返回的並非所想像的那種結果 , 因此在 golang 中要注意這種操作
	for i, v := range "string" {
		fmt.Println(i, v)
	}
}
