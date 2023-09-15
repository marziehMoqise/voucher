package migrations

type Vouchers struct {
	ID          int32  `gorm:"primaryKey"`
	VoucherCode string `gorm:"column:voucherCode;type:varchar(24);unique;not null"`
	GiftAmount  int32  `gorm:"column:giftAmount;not null"`
	Reusability int32  `gorm:"not null"`
	UsedCount   int32  `gorm:"column:usedCount;not null"`
	Type        string `gorm:"type:enum('discount','gift'); not null; index;"`
}
