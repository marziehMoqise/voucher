package voucherSchema

import (
	"apiGolang/utils"
	"github.com/gofiber/fiber/v2"
)

func (req *SetRequest) Validate(ctx *fiber.Ctx) (string, error) {
	errMsg, err := utils.ValidateStruct(req)
	if err != nil {
		switch errMsg {
		case "VoucherCode,required":
			return "The voucherCode field is required.", err
		case "VoucherCode,max":
			return "The voucherCode field must not be greater than 24 characters.", err
		case "GiftAmount,required":
			return "The giftAmount field is required.", err
		case "Reusability,required":
			return "The reusability field is required.", err
		}
	}
	return "", nil
}
