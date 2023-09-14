package dataModel

type Transaction struct {
	ID          int64  `gorm:"column:id"`
	UserID      int64  `gorm:"column:userID"`
	Amount      int64  `gorm:"column:amount"`
	Type        string `gorm:"column:type"`
	Description string `gorm:"column:description"`
	Time        int64  `gorm:"column:time"`
}
