package database

import (
	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var connection *gorm.DB

func setConnectionClient() {
	var err error
	URI := os.Getenv("MYSQL")
	if URI == "" {
		panic("MYSQL not found")
	}
	//zap.L().Info("MYSQL url at env ", zap.Any("url:", URI))
	USER := os.Getenv("MYSQL_USER")
	PASSWORD := os.Getenv("MYSQL_PASSWORD")
	DB := os.Getenv("MYSQL_DB")
	if USER == "" || PASSWORD == "" || DB == "" {
		//zap.L().Error("MYSQL param required")
		panic("some MySQL param (MYSQL_USER-MYSQL_PASSWORD-MYSQL_DB) are missed")
	}

	dsn := USER + ":" + PASSWORD + "@tcp(" + URI + ")/" + DB + "?charset=utf8mb4&parseTime=true&loc=Local"
	connection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            os.Getenv("prepareStmt") == "1",
	})
	if err != nil {
		//zap.L().Error("MYSQL connection error", zap.Error(err))
		panic(err)
	}
	_, err = connection.DB()
	if err != nil {
		panic(err)
	}
}

func GetConnection() *gorm.DB {
	if connection == nil {
		setConnectionClient()
	}
	return connection
}

func IsDuplicateKeyErr(err error) bool {
	if e, ok := err.(*mysqlDriver.MySQLError); ok {
		return e.Number == 1062
	}
	return false
}