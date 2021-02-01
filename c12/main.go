package main

import "fmt"

func main() {
	// 額外補充 c11 的 Array 使用
	// 在使用 for 走訪 Array 的時候 , 修改 Array 的 element 值是沒用的
	arr := [5]int{1, 2, 4, 8, 16}
	for _, e := range arr {
		e = e * e
	}
	for _, e := range arr {
		fmt.Println(e)
	}
	// 如果要修改 Array 的 element , 要以 index 走訪陣列 , 再修改 Array[index] 的值 , 如下
	for i := 0; i < len(arr); i++ {
		// 注意這邊 , 是透過 arr[i] 去修改的
		arr[i] = arr[i] * arr[i]
	}
	for _, e := range arr {
		fmt.Println(e)
	}
	// 以下開始練習 slice , 以下 code 無法正常執行 , 因為宣告 scores 的 slice 長度是 0
	// scores := make([]int, 0, 10)
	// scores[5] = 9033
	// fmt.Println(scores)

	// 但是上面的情況違反理解的 slice 可以彈性調整的功能 , 因此我們可以用以下的寫法
	scores := make([]int, 0, 10)
	scores = scores[0:6]
	scores[5] = 9033
	fmt.Println(scores)

	// 談談 append
	scores2 := make([]int, 0, 5)
	c := cap(scores2)
	fmt.Println(c)
	for i := 0; i < 25; i++ {
		scores2 = append(scores2, i)
		fmt.Println(scores2)
		fmt.Println(c)
		// 如果容量改變 ,  go 會為了能夠放進去這些新 data , 會進行增長 slice 的長度
		if cap(scores2) != c {
			c = cap(scores2)
			fmt.Println(c)
		}
	}
	// 練習範例 , append 會在 slice 的最後一個 element 插入 , 因此這個範例輸出會是 [0 0 0 0 0 9033]
	scores3 := make([]int, 0, 5)
	scores3 = append(scores3, 9332)
	fmt.Println(scores)
}
