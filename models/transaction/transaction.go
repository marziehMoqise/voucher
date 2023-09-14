package transaction

import (
	"apiGolang/database"
	transactionDataModel "apiGolang/models/transaction/dataModel"
)

func Insert(userID, amount int64, transactionType string) error {
	db := database.GetConnection()
	transaction := transactionDataModel.Transaction{
		UserID: userID,
		Amount: amount,
		Type:   transactionType,
		//Time: time.Now(),
	}
	return db.Create(&transaction).Error;
}
