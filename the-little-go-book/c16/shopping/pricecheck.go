package shopping

import "shopping/db"

// Item 商品資料結構 -> 這段在最後新增 models package 後 , 就不存在其意義了
type Item struct {
	Price float64
}

// PriceCheck 檢查 item 是否存在 , 並傳回價值
func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
