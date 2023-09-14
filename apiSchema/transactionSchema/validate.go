package transactionSchema

import (
	validate "apiGolang/services"
	"github.com/gofiber/fiber/v2"
)

func (req *ListRequest) Validate(ctx *fiber.Ctx) (int, error) {
	errMsg, err := validate.Struct(req)
	if err != nil {
		switch errMsg {
		case "mobile,required":
			return 400, err
		case "mobile,max":
			return 400, err
		}
	}

	return 200, nil
}
