package migrations

type Vouchers struct {
	ID          int64  `gorm:"primaryKey"`
	VoucherCode string `gorm:"column:voucherCode;type:varchar(24);unique;not null"`
	GiftAmount  int64  `gorm:"column:giftAmount;not null"`
	Reusability int64  `gorm:"not null"`
	UsedCount   int64  `gorm:"column:usedCount; not null; default:0"`
	Type        string `gorm:"type:enum('discount','gift'); not null; default:gift; index;"`
}
