package main

// Point to value -> 即使你不打算改變數據，也要考慮到創建一个大 struct 複本的開銷。相反，你可能有一個小 struct，例如
type Point struct {
	X int
	Y int
}
