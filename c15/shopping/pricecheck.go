package shopping

import (
	// 作者提到多數人認為 , 已經在 shopping 目錄下 , 導入 package 為何還需要在寫一次 shopping , 實際上其實在導入 package 時 , 是去查找 GOPATH/src 下的 package 名稱
	// 因此這邊的意思是導入 "$GOPAHT/src/shopping/db" 這個包
	"shopping/db"
)

func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
