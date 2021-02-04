package db

// 在 golang 中 , 若是要給外部引用的 package , 在宣告的時候第一個字必須大寫 , 類似 java 的 public 定義

// Item 是一個商品
type Item struct {
	Price float64
}

// LoadItem 是個讀取商品價格的方法
func LoadItem(id int) *Item {
	return &Item{
		Price: 9.99,
	}
}
