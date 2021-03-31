## Reference
- [Medium部落客整理Go學習資源](https://medium.com/@john.lin/golang-%E5%AD%B8%E7%BF%92%E7%AD%86%E8%A8%98-582cad359738)
- [Go 101 ebook](https://gfw.go101.org/article/101.html)
- [The little go book](https://www.openmymind.net/assets/go/go.pdf)
- [The little go book 中文版](https://github.com/songleo/the-little-go-book_ZH_CN)
- [An Intro programming in go ebook](http://www.golang-book.com/books/intro)
- [Awesome Go github](https://github.com/avelino/awesome-go#web-frameworks)
- [tutorial edge](https://tutorialedge.net/golang/go-basic-types-tutorial/)

## :star: 基礎概念篇

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

#### Tips
:star: 關於 golang 中的 func 語法 
:star: go 如何定義 method : A method is a function with a special receiver argument.
```
func (reciever:繼承的方法) funcName(參數)(回傳值型態) { ... }
```

### c05
學習 `struct` , 雖然 go 不像 java / c# / c++ 等語言有 OO 的特性 , 但是在運用 go 的 struct 其實還是有類似 java 在宣告 class 的影子 , 範例定義了一個賽亞人的 struct
![struct definition](/assets/struct_def.png)

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
一個問題 , 在真正實做的時候 , 該選擇使用 value 還是用 pointer to value ? 書中提到多數情況下 , 在不清楚自己要用甚麼的時候 , 優先選用 pointer to value  , 書中提到以下幾種情況 , 直接使用 pointer to value
- 宣告 local variable 賦值
- struct 的 field
- 在 function 中返回值的時候
- 傳遞 function 參數的時候
- 方法的 receiver

#### Tips
書中提到 , 用 value 傳遞的方式是一種確保 data 不可變的好方法 , 有些時候很好用 , 但有些時候這並不是你所想的那樣 , 而為什麼基本上都用 pointer to value 的做法呢? , 書中說明 , 即使不考慮會去改變 data 的值 , 但還是得考慮到當這個 data 的 struct 非常大的時候 , 要複製這個 struct 遠比複製該 struct 的 address 來的困難許多 , 因此除非很確定這個 struct 非常小 , 不然基本上都是 pointer to value 的使用情況較多

### c11
Array , 和大多數語言一樣 , golang 的 Array 並不能動態的變大變小 , 因此實用性不高(雖然效能好) , 在多數情況下 , 我們會提前處理 data , 這種情況並不會提前知道 array 有多大 , 因此 , 在如果要能動態改變 Array 大小 , 要使用 slice (切片)

### c12
在使用 go 中 , 一般較少使用 Array , 多數情況下會使用 slice , [額外參考](https://michaelchen.tech/golang-programming/array-slice/) , 練習 slice 操作 , 以及對於 append 和 cap 的使用進行範例練習
聊聊 slice 宣告的時候注意的地方
![About slice len/cap](/assets/len_vs_cap.png)

### c13
聊聊 map , 在其他語言裡 , 被稱為 dict 或是 hashmap , 定義一組 key:value 的 collection . map 跟 slice 一樣 , 可以用 `make` 創建 , 範例中練習 map 的操作方式

### c14
回到一個問題 , 在 c10 的時候 , 談論到該用 value 還是 pointer to value , 在這裡 slice/map , 在提出討論一次 . 究竟是用 `a := make([]Saiyan, 10)` 還是 `b := make([]*Saiyan, 10)` 呢 , 大部分 developer 認為在 func 中傳遞 b 或返回 b 會比 a 效率來的好 , 然而 , 這裡的傳遞或返回都是一個 slice 的複本 , 這本身就是一個引用 , 因此其實沒有區別 . 真正的區別在於 , 當開始去改變這個 slice 或 map 的值的時候 , 在這點上跟 c10 提到的一樣 , 當這個 "值" 的資料結構大的時候 , pointer to value 會比 value 來的好 , 因此 , 需要探討的是 , 這個 "值" 的定義影響選擇 , 而不是在使不使用 slice 或 map 的時候

#### Tips
在繼續往下之前 , 書中提到 , 上述的內容適用多數的情況 , 但實際上還有一些非常極端的例子 , 這些例子並不常見 , 但作者希望能夠透過書中打下的基礎讓讀者能夠去思考與應付這些情況

## :star: 進階應用篇
書中的 chapter 4 開始 , topic 為 "Code Organization and Interfaces"

### c15
在開發大型專案的時候 , 為了便於管理 , 我們往往會切分不同的檔案並進行引用 , 因此必須認識在 golang 裡面的 package . 現在開始 , 書中作者希望帶領讀者開始學習如何組織自己的程式碼結構 , 並且學習操作 golang 的 package , 與解釋 package 的概念

書中假設 , 如果要開發一個購物系統 , 可能需要從創建一個 shopping 目錄開始 , 並且將這個目錄放到 `$GOPATH/src` 下 (早期 golang 在還沒有套件管理的時候 , 確實是這樣 , 後來有了套件管理的工具 , 已經可以在非 `$GOPATH/src` 的路徑下進行開發了)

接著 , 創建這樣一個 shopping 目錄(並在該目錄執行 `go mod init shopping` , 可以試試如果沒有執行的情況下運行程式 `go run` 會不會有問題) , 但其實並不希望將所有檔案都塞進這個資料夾 , 因此 , 在這裡假設開發這套系統中 , 需要撰寫一些操作 database 的邏輯 , 故會在 shopping 目錄下在創建一個名為 db 的目錄 , 然後才開始放入一些像 `db.go` 的文件(裡面撰寫著操作 db 的邏輯 , 範例中並沒有真的去連線 db , 作者僅是為了展示如何組織程式碼)

再來 , 在 shopping 目錄下創建一個 `pricecheck.go` 的檔案(裡面撰寫對於購物系統檢查價格的邏輯) , 如果只是單純要撰寫被引用的 package , 那麼到這裡就是完整的創建一個可以被引用的 package 了 , 但如果要執行一個程式去引用剛剛創建的這些 package , 還需要一個 main func , 因此最後

作者的習慣是在 shopping 目錄下在創建一個 main 目錄 , 然後在 main 目錄下創建一個 `main.go` 檔案 , 至此 , c15 範例練習完成

#### Tips
注意書裡面所撰寫的內容部分有部分已經過時 , 例如該篇的 package 的套件管理方式 , 書中提到的套件管理工具目前已非主流 , 而 golang 官方也在後面發布了自己預設的套件管理工具 go module . 因此該篇將主要放在理解概念 . 過程中展示的 `go mod init shopping` 就是在引入 go module 這個套件管理工具 , 詳細可參考[官網](https://golang.org/doc/tutorial/create-module)

### c16
在 c15 內容聊到了 package 的 import , c16 將深入聊聊一種情況 "Cyclical Imports" 又稱為循環導入(循環依賴) , 這種情況發生在當 package A 引用了 package B , 且 package B 也"直接或間接"引用了 package A . c16 以 c15 的程式結構展示該情況的發生作為範例

首先將 shopping/db/db.go 中的 `type Item struct` 區段搬到 shopping/pricecheck.go 然後執行 `go run main/main.go` , 你會得到一些錯誤
```
db/db.go:9:24: undefined: Item
db/db.go:10:10: undefined: Item
```
接著 , 修改 shopping/db/db.go 將內容改為 , 然後再次執行
```
package db

import (
  "shopping"
)

func LoadItem(id int) *shopping.Item {
  return &shopping.Item{
    Price: 9.001,
  }
}
```
你會得到
```
package command-line-arguments
        imports shopping
        imports shopping/db
        imports shopping: import cycle not allowed
```
為了解決這樣的問題 , 我們新定義了一個 package "models" , 而這個 models 裡面宣告了剛剛這個共享的 Item struct , 因此最後的目錄結構會變成這樣
```
- shopping
    pricecheck.go
    - db
        db.go
    - models
        item.go
    - main
        main.go
```
在平常撰寫 go 中 , 會經常定義類似這樣的 struct , 甚至可能還會命名一個 utilities 的 package 來定義 struct 等等 , 但請注意 , 最重要的原則是在定義這類的 struct 時 , 千萬不要引用在 shopping 或 shopping 下的任何 package , 在後面 , 會看到一些例子使用 `interface` 來解決像這類的依賴關係

### c16-2
這篇小章節來聊聊針對 c16 用到的觀念卻沒有說明到的地方.
在 golang 中 , go 用大小寫來定義了 visibility (可見性) 的宣告 , 舉例來說 , 當一個 package A 導入了 package B , 那個 B 中的 func 或 variables 甚至是 struct 等等 , 宣告時的第一個字都需要是大寫 , 否則 package A 是無法調用的 . 看看下面的情況
```
package B

// 試試看把 NewItem 換成 newItem , 你會得到一些編譯錯誤
func NewItem() *Item {
    ...    
}
```
上面情況同理也適用在 struct 
```
// 試試看把 Item 換成 item
type Item struct {
    // 試試看把 Price 換成 price , 其他 package 會無法調用
    Price float64
}
```
因此 , 結論是 , 小寫字母開頭的宣告下 , 只有在同一個 package 可以調用 , 而如果要給不同的 package 調用 , 宣告時的第一個字母必須使用大寫

#### Tips
書中還有提到 package management 和如何導入 third-party 的 package 操作方式 , 這邊不細探 , 未來會有很多機會用到

### c17
談談在 golang 裡面的 interface , 在 golang 裡 interface 是一種型別 , 是定義 func 的 collection , 它僅定義 , 不實作 . 相當程度的使用 interface 可以從程式碼實作中解耦一些行為 , :star: 目前在這邊讀 interface 還有些無法理解 interface 的應用 , 之後有機會再回過頭來看一遍

這邊先以官方的一個 interface 練習來感受一下 interface 的操作和用法 [官方範例](https://gobyexample.com/interfaces) 

### c18
OK , 接下來開始聊聊 error handling (錯誤處理) . Go 在進行錯誤處理通常傾向使用 "返回值(return value)" 進行處理 , 並非用 "例外處理(exceptions)" 範例中展示 `strconv.Atoi` 功能為 將 string 轉換成 integer 型別的 function 的錯誤處理案例

#### Tips
golang 也支持自定義 error 的 interface , 但是必須要實作 golang 內建的 `error` interface
```
type error interface {
    Error() string
}
```
在普遍的情況下 , 可以自己定義 error 透過 import "error" 的 package , 並使用它內部的 New function
```
import (
  "errors"
)
func process(count int) error {
  if count < 1 {
    return errors.New("Invalid count")
  }
  ...
  return nil
}
```
最後補充一點 , go 語言中也有 `panic` 和 `recover` 的 function , 但很少使用

### c19
談完了 error handling , 接下來更深入探討發生錯誤的情況 , 常見的一些問題 . c19 主要聊聊 golang 中的 defer . 雖然 golang 有 garbage collector , 但設想一種情形 , 一段運行讀檔操作檔案的程式碼原本運行的好好的 , 但突然某一個環節出了錯誤 , 這個錯誤最終導致這段操作檔案的程式碼 "沒有釋放(沒有呼叫 Close 函數)" 掉 "讀檔" 這個 resource , 在這種情況下 , 資源用完並沒有被回收 . 為了避免這種情況發生 , golang 提供了 `defer` 這個關鍵字 , 它的原理是 , 會在 "函數返回之前執行 , 並且以 LIFO 後進先出的方式執行" , 接著讓我們透過 c19 展示 `defer` 的使用方式

#### Tips
c19 的 defer.go 展示 defer 的工作原理 , 運行後可以查看出差異在哪

### c20
空 interface 和 型別轉換 (conversions) , 這章先跳過 , 這章有點難理解 ~"~

### c21
String 與 Byte Array 是有著密切關係的 , 在 c21 範例中 , 可以看出 string 和 byte array 之間能夠輕易在之間做轉換 . 而實際上 , 其他型別的轉換原理也是類似這樣的做法 , 像是 int64() 或 int32() 這類 function .
而在使用 byte 和 string 的時候 , 需要注意的是 , 我們經常會使用到 string 或 byte 型別的操作 , 但實際上我們創建的都是 data 的複本(copy) , :star: 這是因為 string 和 byte 的 "immutable 不可變性"

c21 範例中也說明 , 當遇到 string 中含有 unicode 中 runes 這種特殊字元時 , 在計算該 string 的長度可能不會得到正確的答案 , 如同 c21 範例所示 . 而如果在使用 for 迴圈 iterat 一個 string 的時候 , 返回的結果也不是一般所想的那樣  

### c22
使用 type func , func 是第一種 class 型別 , 這種型別可以用在任何地方 , 例如 struct , 適當的運用 type func 可以有效的降低程式碼的耦合度(解耦一些程式碼邏輯)

## 併發
:star: 暫時不讀 , 之後再回來讀