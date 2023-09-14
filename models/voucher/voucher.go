package voucher

import (
	"apiGolang/database"
	voucherDataModel "apiGolang/models/voucher/dataModel"
	"gorm.io/gorm"
)

func GetVoucherByCode(voucherCode, voucherType string) (voucher voucherDataModel.Voucher, err error) {
	db := database.GetConnection()
	err = db.Take(&voucher, "voucherCode = ? AND type = ?", voucherCode, voucherType).Error

	return voucher, err
}

func IncreaseUsedCount(voucherID int64) error {
	db := database.GetConnection()
	return db.Exec("UPDATE vouchers SET usedCount = ? WHERE id = ?", gorm.Expr("usedCount + ?", 1), voucherID).Error
}
