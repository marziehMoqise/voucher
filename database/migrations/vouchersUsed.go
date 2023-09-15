package migrations

func (VouchersUsed) TableName() string {
	return "vouchersUsed"
}

type VouchersUsed struct {
	ID        int32 `gorm:"primaryKey"`
	UserID    int32 `gorm:"column:UserID;not null; index"`
	VoucherID int32 `gorm:"column:voucherID;not null; index"`
	Time      int32 `gorm:"not null"`
}
