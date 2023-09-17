package voucherUsedSchema

import (
	"apiGolang/utils"
	"github.com/gofiber/fiber/v2"
)

func (req *ListRequest) Validate(ctx *fiber.Ctx) (string, error) {
	if len(req.Mobile) > 0 {
		isValid, mobile := utils.MobileValidate(req.Mobile)
		if isValid {
			req.Mobile = mobile
		}
	}

	return "0", nil
}
