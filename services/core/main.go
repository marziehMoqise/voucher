package main

import (
	"apiGolang/controllers/transaction"
	"apiGolang/controllers/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app:=fiber.New()
	app.Use(logger.New())

	app.Post("/user/gift", user.Gift)
	app.Post("/transactions", transaction.List)

	if err := app.Listen(":7575"); err != nil {
		//panic(err)
	}
}
