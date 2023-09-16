package migrations

func (VouchersUsed) TableName() string {
	return "vouchersUsed"
}

type VouchersUsed struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64 `gorm:"column:UserID;not null; index"`
	VoucherID int64 `gorm:"column:voucherID;not null; index"`
	Time      int64 `gorm:"not null"`
}
