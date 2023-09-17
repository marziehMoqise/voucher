package user

import (
	"apiGolang/apiSchema/userSchema"
	userModel "apiGolang/models/user"
	"apiGolang/utils"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetBalance(ctx *fiber.Ctx) error {
	req := new(userSchema.BalanceRequest)
	ctx.BodyParser(req)

	errCode, err := req.Validate(ctx)
	if err != nil {
		return utils.ResponseError(ctx, errCode)
	}

	user, err := userModel.GetByMobile(req.Mobile)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ResponseError(ctx, "user not found")
		}

		log.Error("Get user by mobile", zap.Error(err))
		return utils.ResponseError(ctx, "operation failed(20180)")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"balance": user.Balance,
		},
	})
}
