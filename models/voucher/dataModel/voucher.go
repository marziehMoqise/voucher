package dataModel

type Voucher struct {
	ID          int64  `gorm:"column:id"`
	VoucherCode string `gorm:"column:voucherCode"`
	Reusability int64  `gorm:"column:reusability"`
	UsedCount   int64  `gorm:"column:usedCount"`
	Type        string `gorm:"column:type"`
	GiftAmount  int64  `gorm:"column:giftAmount"`
}
