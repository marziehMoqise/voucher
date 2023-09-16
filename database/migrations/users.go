package migrations

type Users struct {
	ID      int64  `gorm:"primaryKey"`
	Mobile  string `gorm:"type:varchar(20);unique;not null"`
	Balance int64  `gorm:"not null"`
}
