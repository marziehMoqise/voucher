package transactionSchema

import (
	"apiGolang/utils"
	"github.com/gofiber/fiber/v2"
)

func (req *ListRequest) Validate(ctx *fiber.Ctx) (string, error) {
	errMsg, err := utils.ValidateStruct(req)
	if err != nil {
		switch errMsg {
		case "Mobile,required":
			return "The mobile field is required.", err
		case "Mobile,max":
			return "The mobile field must not be greater than 20 characters.", err
		}
	}

	isValid, mobile := utils.MobileValidate(req.Mobile)
	if isValid {
		req.Mobile = mobile
	}

	return "", nil
}
