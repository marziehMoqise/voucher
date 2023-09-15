package migrations

type Users struct {
	ID      int32  `gorm:"primaryKey"`
	Mobile  string `gorm:"type:varchar(20);unique;not null"`
	Balance int32  `gorm:"not null"`
}
