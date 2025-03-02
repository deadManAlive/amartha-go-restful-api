package domain

type Product struct {
	ProductID   uint64   `gorm:"primaryKey;column:id"`
	Name        string   `gorm:"column:name;length:255"`
	Description string   `gorm:"column:description;length:255"`
	Price       float64  `gorm:"column:price"`
	StockQty    int      `gorm:"column:stock_qty"`
	CategoryId  uint64   `gorm:"column:category_id"`
	SKU         string   `gorm:"column:sku"`
	TaxRate     float64  `gorm:"column:tax_rate"`
	Category    Category `gorm:"foreignKey:CategoryId;references:Id"`
}
