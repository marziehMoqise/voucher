package transaction

import (
	"apiGolang/apiSchema/transactionSchema"
	"apiGolang/models/transaction"
	userModel "apiGolang/models/user"
	"apiGolang/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func List(ctx *fiber.Ctx) error {

	req := new(transactionSchema.ListRequest)
	ctx.BodyParser(req)

	errCode, err := req.Validate(ctx)
	if err != nil {
		return utils.ResponseError(ctx, errCode)
	}

	user, err := userModel.GetByMobile(req.Mobile)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return utils.ResponseError(ctx, "user not found")
		}

		log.Error("Get user by mobile", zap.Error(err))
		return utils.ResponseError(ctx, "operation failed(20160)")
	}

	transactions, err := transaction.GetByUserID(user.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error("Get transaction by userID", zap.Error(err))
		return utils.ResponseError(ctx, "operation failed(20161)")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"transactions": transactions,
		},
	})
}