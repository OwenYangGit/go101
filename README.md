## Reference
- [Medium部落客整理Go學習資源](https://medium.com/@john.lin/golang-%E5%AD%B8%E7%BF%92%E7%AD%86%E8%A8%98-582cad359738)
- [Go 101 ebook](https://gfw.go101.org/article/101.html)
- [The little go book](https://www.openmymind.net/assets/go/go.pdf)
- [The little go book 中文版](https://github.com/songleo/the-little-go-book_ZH_CN)
- [An Intro programming in go ebook](http://www.golang-book.com/books/intro)
- [Awesome Go github](https://github.com/avelino/awesome-go#web-frameworks)

### c01 
雖然僅僅展示了簡單的 println , 但這邊需要注意的是 , 在 main package 下運行的程式碼 , 注意到這個 main , 它作為運行一個程式碼的入口點 , 可以的話能嘗試將 package 的 main 改成別的名稱 , 在運行時會得到一個錯誤 , 而將 func 的 main 改成別名稱 , 運行時會得到另一種錯誤

### c02
go 的包 import , 這邊開始學第一個 "關鍵字" (import) , 通常在使用 go 開發時 , 會引用外部的第三方 package , 因此 import 的功能非常重要 , 在 c02 範例中 , 引用了 go 的內置套件 fmt 與 os , 後面會在學習到如何引用第三方的"外部"套件 , 這邊比較特別的是 , go 的 import 非常嚴格 , 如果 import 了套件在程式碼沒有使用 , 那個編譯就不會 , 這點需要注意

#### tips
如果想要看 go 的 fmt 相關文件可以透過 https://golang.org/pkg/fmt/#Println , 如果不能上網 , 也可以在本地運行 godoc -http=:6060 , 接著打開 http://localhost:6060

### c03
從範例來看 go 在聲明變數時 , 它的運作方式為 var power int (先聲明一個變數 power 型別為 int 且賦予值 0) , power = 9000 (宣告 power 重新賦值為 9000) , 範例中也提出第二種簡潔的寫法 `:=` , 這個做法會讓 go 在第一次運行時自動判別變數型別是甚麼 , 並執行賦值的動作 , 需要注意的是 , 這只有在第一次宣告該變數時才能這樣寫 , 之後若要重新賦值仍是使用 `=` , 因此這邊要能清楚了解到使用 `:=` 與 `=` 的差別在哪

#### tips
在 c03 最後 , 需要特別注意的是 , 變數和 package 的情況類似 , 當你宣告了一個變數 name , 但在程式中並沒有使用它 , 那麼編譯時將會拋出錯誤

### c04
這個範例學習到的是如何聲明 "func" , go 的 func 支援多值返回 , 範例中可以練習如何取回 func 的返回值 , 以及一些常見的處理手法 , 較特別的是 , ` _ ` 符號是空白標識符 , 常見的情況下可以用 ` _ ` 來丟棄一個返回值 , 而如果 func 返回的值類型都一樣那就可以直接用下面的方式來做聲明
```
func add(a,b int) int {

}
```

### c05
學習 `struct` , 雖然 go 不像 java / c# / c++ 等語言有 OO 的特性 , 但是在運用 go 的 struct 其實還是有類似 java 在宣告 class 的影子 , 範例定義了一個賽亞人的 struct

### c06
代入指針概念 , 了解變數跟賦值與指針之間的操作 , c06 範例展示操作 `struct` 的一些操作手法 . 並演示指針相關的概念 , 以及自己一些操作嘗試
![Pointer](/assets/pointer_in_go.png)

### c07
練習 `struct` 與一個 function 關聯的操作方式

### c08
`struct` 裡面可以有任何 type 的宣告 , 可以是整數 , 字串 , 甚至是 `struct` , 範例展示 struct 裡面還有 struct

### c09
`struct` 裡面引用其他 `struct` , 有點像繼承的概念 , 範例展示 go 在這種情況內的做法

### c10
一個問題 , 在真正實做的時候 , 該選擇使用 value 還是用 pointer to value ? 書中提到多數情況下 , 在不清楚自己要用甚麼的時候 , 優先選用 pointer to value  , 而下面情況下 , 都是屬於用 pointer to value
- 宣告 local variable 賦值
- struct 的 field
- 在 function 中返回值的時候
- function 的參數
- 方法的 receiver
