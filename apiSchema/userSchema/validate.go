package userSchema

import (
	validate "apiGolang/services"
	"github.com/gofiber/fiber/v2"
)

func (req *GiftRequest) Validate(ctx *fiber.Ctx) (int, error) {
	errMsg, err := validate.Struct(req)
	if err != nil {
		switch errMsg {
		case "voucherCode,required":
			return 400, err
		case "mobile,required":
			return 400, err
		case "mobile,max":
			return 400, err
		}
	}

	return 200, nil
}
