package main

import (
	"fmt"
)

func main() {
	var power int
	power = 9000
	fmt.Printf("It's over %d\n", power)
	// 另一種寫法
	power2 := getPower()
	fmt.Printf("It's over %d\n", power2)
}

func getPower() int {
	return 9001
}
