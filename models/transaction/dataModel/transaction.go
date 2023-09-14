package dataModel

//import "time"

type Transaction struct {
	ID     int64  `gorm:"column:id"`
	UserID int64  `gorm:"column:userID"`
	Amount int64  `gorm:"column:amount"`
	Type   string `gorm:"column:type"`
	//Time   time.Time `gorm:"column:time"`
}
