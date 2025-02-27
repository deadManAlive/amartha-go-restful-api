package domain

type Category struct {
	Id   uint64 `gorm:"primary_key; column:id"`
	Name string `gorm:"column:name"`
}
