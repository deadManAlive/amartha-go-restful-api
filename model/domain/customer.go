package domain

type Customer struct {
	Id         uint64 `gorm:"primary_key;column:id"`
	Name       string `gorm:"column:name;type:varchar(100);"`
	Email      string `gorm:"column:email;type:varchar(255);"`
	Phone      string `gorm:"column:phone;type:varchar(20);"`
	Address    string `gorm:"column:address;type:varchar(255);"`
	LoyaltyPts int    `gorm:"column:loyalty_pts;type:int(11);"`
}
