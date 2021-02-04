package main

import (
	"fmt"
	"math"
)

// Interface 是一個被命名為 collection of method 的定義

// 現實世界的理解是 geometry 是定義 "幾何圖形" 的介面 (這個介面定義了只要屬於 "幾何圖形" 就應該會有的 method)
// geometry : 名為 geometry 的 interface 是 area() 和 perim 方法(method) 的 collection
type geometry interface {
	area() float64  // 面積
	perim() float64 // 周長
}

// rect 定義長方形物件的 struct
type rect struct {
	width, height float64
}

// circle 定義圓形物件的 struct
type circle struct {
	radius float64
}

// area 算長方形面積的方法
func (r rect) area() float64 {
	// 長方形的面積算法 = 長(高) * 寬
	return r.width * r.height
}

// perim 算長方形周長的方法
func (r rect) perim() float64 {
	// 長方形的周長算法 = 2個長 + 2個寬
	return r.width*2 + r.height*2
}

// area 算圓形面積的方法
func (c circle) area() float64 {
	// 圓形的面積算法 = 圓周率 * 半徑 * 半徑
	return math.Pi * c.radius * c.radius
}

// perim 算圓形周長的算法
func (c circle) perim() float64 {
	// 圓形的周長算法 = 直徑 * 圓周率 = 2 倍半徑 * 圓周率
	return 2 * math.Pi * c.radius
}

// 如果一個變數存放的是一個 interface , 如下 , g 變數存放了 geometry 這個 interface
// 那麼 g 這個變數就能呼叫 geometry 這個 interface 裡面定義的 method
// measure 這個普通的 function , 它能作用於所有代入有"實作' geometry 的物件
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
// 範例的註解表示因為 circle 和 rect 的 struct 都實做了 geometry 因此我們可以將由這兩種 struct 初始化出來的 instance 當作參數傳入到 measure 的 function

func main() {
	r := rect{
		width:  3,
		height: 4,
	}
	c := circle{
		radius: 5,
	}

	measure(r)
	measure(c)
}
