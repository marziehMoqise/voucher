package transaction

import (
	"apiGolang/apiSchema/transactionSchema"
	"apiGolang/models/transaction"
	userModel "apiGolang/models/user"
	response "apiGolang/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func List(ctx *fiber.Ctx) error {

	req := new(transactionSchema.ListRequest)
	ctx.BodyParser(req)

	user, err := userModel.GetByMobile(req.Mobile)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return response.ResponseError(ctx, "user not found")
		}
		return response.ResponseError(ctx, "operation failed(20160)")
	}

	transactions, err := transaction.GetByUserID(user.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return response.ResponseError(ctx, "operation failed(20161)")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"transactions": transactions,
		},
	})
}