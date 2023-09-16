package voucher

import (
	"apiGolang/apiSchema/voucherSchema"
	"apiGolang/database"
	"apiGolang/models/voucher"
	voucherDataModel "apiGolang/models/voucher/dataModel"
	"apiGolang/utils"
	"github.com/gofiber/fiber/v2"
)

func Set(ctx *fiber.Ctx) error {

	req := new(voucherSchema.SetRequest)
	ctx.BodyParser(req)

	errCode, err := req.Validate(ctx)
	if err != nil {
		return utils.ResponseError(ctx, errCode)
	}

	err = voucher.Create(&voucherDataModel.Voucher{
		VoucherCode: req.VoucherCode,
		Reusability: req.Reusability,
		Type:        "gift",
		GiftAmount:  req.GiftAmount,
	})

	if err != nil {
		if database.IsDuplicateKeyErr(err) {
			return utils.ResponseError(ctx, "The voucherCode has already been taken.")
		}
		return utils.ResponseError(ctx, "The operation failed")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"message": "The operation was successful",
		},
	})
}