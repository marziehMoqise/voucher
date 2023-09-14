package transaction

import (
	"apiGolang/apiSchema/transactionSchema"
	"apiGolang/database"
	transactionDataModel "apiGolang/models/transaction/dataModel"
	"time"
)

func Insert(userID, amount int64, transactionType, description string) error {
	db := database.GetConnection()
	transaction := transactionDataModel.Transaction{
		UserID:      userID,
		Amount:      amount,
		Type:        transactionType,
		Description: description,
		Time:        time.Now().Unix(),
	}
	return db.Create(&transaction).Error
}

func GetByUserID(userID int64) (transactions []transactionSchema.ListResponse, err error) {

	db := database.GetConnection()
	err = db.Model(transactionDataModel.Transaction{}).Find(&transactions, "userID = ?", userID).Error

	return transactions, err
}
