package migrations

type Transactions struct {
	ID          int64  `gorm:"primaryKey"`
	UserID      int64  `gorm:"column:UserID;not null; index"`
	Amount      int64  `gorm:"not null"`
	Type        string `gorm:"type:enum('increase', 'decrease'); not null; index;"`
	Description string `gorm:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Time        int64  `gorm:"not null"`
}
