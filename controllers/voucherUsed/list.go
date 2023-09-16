package voucherUsed

import (
	voucherUsedSchema "apiGolang/apiSchema/voucherUsed"
	userModel "apiGolang/models/user"
	voucherModel "apiGolang/models/voucher"
	"apiGolang/models/voucherUsed"
	"apiGolang/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func List(ctx *fiber.Ctx) error {
	req := new(voucherUsedSchema.ListRequest)
	ctx.BodyParser(req)

	errCode, err := req.Validate(ctx)
	if err != nil {
		return utils.ResponseError(ctx, errCode)
	}

	var userID, voucherID int64
	if req.Mobile != "" {
		user, err := userModel.GetByMobile(req.Mobile)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return utils.ResponseError(ctx, "user not found")
			}

			log.Error("Get user by mobile", zap.Error(err))
			return utils.ResponseError(ctx, "operation failed(20170)")
		}
		userID = user.ID
	}

	if req.VoucherCode != "" {
		voucher, err := voucherModel.GetVoucherByCode(req.VoucherCode, "gift")
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return utils.ResponseError(ctx, "voucherCode not found")
			}

			log.Error("Get voucher by voucherCode", zap.Error(err))
			return utils.ResponseError(ctx, "operation failed(20171)")
		}
		voucherID = voucher.ID
	}

	vouchersUsed, err := voucherUsed.List(userID, voucherID)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error("Get list of voucher used", zap.Error(err))
		return utils.ResponseError(ctx, "operation failed(20172)")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"vouchersUsed": vouchersUsed,
		},
	})
}
