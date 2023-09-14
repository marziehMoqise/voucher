package voucherUsed

import (
	voucherUsedSchema "apiGolang/apiSchema/voucherUsed"
	"apiGolang/database"
	voucherUsedDataModel "apiGolang/models/voucherUsed/dataModel"
	"gorm.io/gorm"
	"time"
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
		Time:      time.Now().Unix(),
	}
	return db.Create(&voucherUsed).Error
}

func List(userID, voucherID int64) (vouchersUsed []voucherUsedSchema.ListResponse, err error) {

	db := database.GetConnection()

	query := make(map[string]interface{})
	if userID != 0 {
		query["userID"] = userID
	}
	if voucherID != 0 {
		query["voucherID"] = voucherID
	}

	result := db.Table("vouchersUsed").Select("vouchersUsed.*, vouchers.voucherCode as VoucherCode, users.mobile")
	result.Joins("join users on vouchersUsed.userID = users.id")
	result.Joins("join vouchers on vouchersUsed.voucherID = vouchers.id")
	result.Find(&vouchersUsed, query)

	return vouchersUsed, result.Error
}
