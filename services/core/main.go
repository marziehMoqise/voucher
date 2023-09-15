package main

import (
	"apiGolang/controllers/transaction"
	"apiGolang/controllers/user"
	"apiGolang/controllers/voucherUsed"
	"apiGolang/database"
	"apiGolang/database/migrations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	db := database.GetConnection()
	db.AutoMigrate(&migrations.Users{})
	db.AutoMigrate(&migrations.Transactions{})
	db.AutoMigrate(&migrations.Vouchers{})
	db.AutoMigrate(&migrations.VouchersUsed{})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/user/gift", user.Gift)
	app.Post("/transactions", transaction.List)
	app.Post("/vouchersUsed", voucherUsed.List)

	if err := app.Listen(":7575"); err != nil {
		panic(err)
	}
}
