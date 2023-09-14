package dataModel

type User struct {
	ID      int64  `gorm:"column:id"`
	Mobile  string `gorm:"column:mobile"`
	Balance int64  `gorm:"column:balance"`
}
