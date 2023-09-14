package voucherUsed

import (
	"apiGolang/database"
	voucherUsedDataModel "apiGolang/models/voucherUsed/dataModel"
	"gorm.io/gorm"
)

func GetVoucherUsedByUserID(voucherID, userID int64) (result *gorm.DB, err error) {
	db := database.GetConnection()
	var voucherUsed voucherUsedDataModel.VoucherUsed
	result = db.Take(&voucherUsed, "voucherID = ? AND userID = ? ", voucherID, userID)

	return result, result.Error
}


func Insert(userID, voucherID int64) error {
	db := database.GetConnection()
	voucherUsed := voucherUsedDataModel.VoucherUsed{
		VoucherID: voucherID,
		UserID:    userID,
		//Time: time.Now(),
	}
	return db.Create(&voucherUsed).Error;
}
