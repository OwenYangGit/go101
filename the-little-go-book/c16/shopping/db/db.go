package db

import (
	"shopping/models"
)

// // Item 商品資料結構
// type Item struct {
// 	Price float64
// }

// LoadItem 返回一個指針型別的 Item
func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.001,
	}
}
