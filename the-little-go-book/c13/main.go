package main

import "fmt"

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9000
	power, exists := lookup["vegeta"]
	// print 0 和 false , 0 為 int type 的預設整數值
	fmt.Println(power, exists)

	// 使用 len 可以取得 map 中的 key 個數
	total := len(lookup)
	fmt.Println(total)
	// 使用 delete 可以刪除 map 中的某個 key:value
	fmt.Println(lookup)
	delete(lookup, "goku")
	fmt.Println(lookup)

	// map 是可以彈性調整的 , 我們也可以透過 make 傳遞第二個參數設定 map 的初始大小
	// lookup2 := make(map[string]int, 100)

	// 當要在 struct 中的 field 使用 map 時
	type Saiyan struct {
		Name    string
		Friends map[string]*Saiyan
	}

	// 初始化上面的 struct
	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
	goku.Friends["krillin"] = &Saiyan{
		Name: "Krillin", // 克林沒朋友 QQ
	}
	fmt.Println(goku)
	fmt.Println(goku.Friends["krillin"].Name)

	// 另一種初始化的方式
	lookup3 := map[string]int{
		"goku":  9001,
		"gohan": 13001,
	}

	// 練習 for 循環 , 這邊需要注意的是 , 編譯器並不會依照 "goku": 9001 / "gohan": 13001 的宣告順序
	// 也就是說每次編譯過後的進行 for 遍歷後 , 返回的 key:value 並不是固定的 (map 是非有序的)
	for key, value := range lookup3 {
		fmt.Println(key)
		fmt.Println(value)
	}
}
