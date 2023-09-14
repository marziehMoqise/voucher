package voucherUsed

import (
	voucherUsedSchema "apiGolang/apiSchema/voucherUsed"
	userModel "apiGolang/models/user"
	voucherModel "apiGolang/models/voucher"
	"apiGolang/models/voucherUsed"
	response "apiGolang/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func List(ctx *fiber.Ctx) error {
	req := new(voucherUsedSchema.ListRequest)
	ctx.BodyParser(req)

	var userID, voucherID int64
	if req.Mobile != "" {
		user, err := userModel.GetByMobile(req.Mobile)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return response.ResponseError(ctx, "user not found")
			}
			return response.ResponseError(ctx, "operation failed(20170)")
		}
		userID = user.ID
	}

	if req.VoucherCode != "" {
		voucher, err := voucherModel.GetVoucherByCode(req.VoucherCode, "gift")
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return response.ResponseError(ctx, "voucherCode not found")
			}
			return response.ResponseError(ctx, "operation failed(20171)")
		}
		voucherID = voucher.ID
	}

	vouchersUsed, err := voucherUsed.List(userID, voucherID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return response.ResponseError(ctx, "operation failed(20172)")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"vouchersUsed": vouchersUsed,
		},
	})
}
