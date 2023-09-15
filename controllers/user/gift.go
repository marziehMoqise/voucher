package user

import (
	"apiGolang/apiSchema/userSchema"
	"apiGolang/database"
	"apiGolang/models/transaction"
	userModel "apiGolang/models/user"
	voucherModel "apiGolang/models/voucher"
	"apiGolang/models/voucherUsed"
	response "apiGolang/services"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func Gift(ctx *fiber.Ctx) error {

	db := database.GetConnection()

	req := new(userSchema.GiftRequest)
	ctx.BodyParser(req)

	user, err := userModel.FirstOrCreateUserByMobile(req.Mobile)
	if err != nil {
		log.Error("FirstOrCreate user by mobile", zap.Error(err))
		return response.ResponseError(ctx, "operation failed(20150)")
	}

	//getVoucherByCode
	voucher, err := voucherModel.GetVoucherByCode(req.VoucherCode, "gift")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ResponseError(ctx, "voucherCode not found")
		}
		log.Error("Get voucher by code", zap.Error(err))
		return response.ResponseError(ctx, "operation failed(20151)")
	}

	//usedVoucherByUser
	result, err := voucherUsed.GetVoucherUsedByUserID(voucher.ID, user.ID)
	if result.RowsAffected > 0 {
		return response.ResponseError(ctx, "You have already used this voucher code")
	}

	//check voucherCode exceeded
	if voucher.UsedCount == voucher.Reusability {
		return response.ResponseError(ctx, "voucherCode has been exceeded")
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db = db.Begin()

			//insertTransaction
			if err = transaction.Insert(user.ID, voucher.GiftAmount, "increase", "افزایش موجودی کیف پول از طریق هدیه"); err != nil {
				db.Rollback()
				return response.ResponseError(ctx, "operation failed(20152)")
			}

			//updateUserBalance
			if err = userModel.UpdateUserBalance(user.ID, voucher.GiftAmount); err != nil {
				db.Rollback()
				return response.ResponseError(ctx, "operation failed(20153)")
			}

			//insert voucherUsed
			if err = voucherUsed.Insert(user.ID, voucher.ID); err != nil {
				db.Rollback()
				return response.ResponseError(ctx, "operation failed(20154)")
			}

			//update usedCount voucher
			if err = voucherModel.IncreaseUsedCount(voucher.ID); err != nil {
				db.Rollback()
				return response.ResponseError(ctx, "operation failed(20155)")
			}
			db.Commit()
		} else {
			log.Error("Get voucher used by userID", zap.Error(err))
			return response.ResponseError(ctx, "operation failed(20156)")
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"message": strings.NewReplacer("{GiftAmount}", strconv.FormatInt(voucher.GiftAmount, 10)).Replace("The amount of {GiftAmount} was added to your wallet"),
		},
	})
}
