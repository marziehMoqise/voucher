package dataModel

func (VoucherUsed *VoucherUsed) TableName() string {
	return "vouchersUsed"
}

type VoucherUsed struct {
	ID        int64 `gorm:"column:id"`
	VoucherID int64 `gorm:"column:voucherID"`
	UserID    int64 `gorm:"column:userID"`
	Time      int64 `gorm:"column:time"`
}
