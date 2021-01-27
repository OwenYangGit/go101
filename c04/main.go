package main

import "fmt"

func log(message string) {
	fmt.Printf("This is log message %d", message)
}

func add(a, b int) int {
	return a + b
}

func power(name string) (int, bool) {
	if name == "devops" {
		return 1, false
	} else {
		return 0, true
	}
}

func main() {
	value, exists := power("goku")
	if exists == false {
		fmt.Printf("value is %d , exists is %t", value, exists)
	} else {
		fmt.Printf("value is %d , exists is %t \n", value, exists)
	}
	// 如果只想要取得其中一個返回值
	// 丟棄了 int 的返回值
	_, exists2 := power("devops")
	if exists2 == false {
		fmt.Printf("Error!!  %t \n", exists2)
	}
	value2, _ := power("devops")
	// 丟棄了 bool 的返回值
	fmt.Printf("Value = %d \n", value2)
}
