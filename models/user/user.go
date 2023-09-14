package user

import (
	"apiGolang/database"
	"apiGolang/models/user/dataModel"
	"gorm.io/gorm"
)

func FirstOrCreateUserByMobile(mobile string) (user dataModel.User, err error) {
	db := database.GetConnection()
	err = db.FirstOrCreate(&user, dataModel.User{Mobile: mobile}).Error

	return user, err
}

func UpdateUserBalance(userID, amount int64) error {
	db := database.GetConnection()
	return db.Exec("UPDATE users SET balance = ? WHERE id = ?", gorm.Expr("balance + ?", amount), userID).Error
}

func GetByMobile(mobile string) (user dataModel.User, err error) {
	db := database.GetConnection()
	err = db.Take(&user, "mobile = ? ", mobile).Error

	return user, err
}
