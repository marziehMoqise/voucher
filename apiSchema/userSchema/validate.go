package userSchema

import (
	"apiGolang/utils"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func (req *GiftRequest) Validate(ctx *fiber.Ctx) (string, error) {
	errMsg, err := utils.ValidateStruct(req)
	if err != nil {
		switch errMsg {
		case "VoucherCode,required":
			return "The voucherCode field is required.", err
		case "Mobile,required":
			return "The mobile field is required.", err
		case "Mobile,max":
			return "The mobile field must not be greater than 20 characters.", err
		}
	}

	isValid, mobile := utils.MobileValidate(req.Mobile)
	if !isValid {
		return "The mobile is not a valid.", errors.New("mobile isn't valid")
	}

	req.Mobile = mobile

	return "", nil
}


func (req *BalanceRequest) Validate(ctx *fiber.Ctx) (string, error) {
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