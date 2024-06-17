package user

import (
	"apiGolang/apiSchema/userSchema"
	"apiGolang/database"
	"apiGolang/models/transaction"
	userModel "apiGolang/models/user"
	"apiGolang/models/user/dataModel"
	voucherModel "apiGolang/models/voucher"
	voucherDataModel "apiGolang/models/voucher/dataModel"
	"apiGolang/models/voucherUsed"
	"apiGolang/utils"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"sync"
)

func Gift(ctx *fiber.Ctx) error {
	var (
		wg sync.WaitGroup
	)
	db := database.GetConnection()

	req := new(userSchema.GiftRequest)
	if err := ctx.BodyParser(req); err != nil {
		return utils.ResponseErrors(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	errMsg, err := req.Validate(ctx)
	if err != nil {
		return utils.ResponseErrors(ctx, fiber.StatusBadRequest, errMsg)
	}

	wg.Add(2)

	var user dataModel.User
	var voucher voucherDataModel.Voucher
	var fetchErr error

	go func() {
		defer wg.Done()
		userData, err := userModel.FirstOrCreateUserByMobile(req.Mobile)
		if err != nil {
			fetchErr = err
			log.Error("Failed to fetch or create user", zap.Error(err))
		}
		user = userData
	}()

	go func() {
		defer wg.Done()
		voucherData, err := voucherModel.GetVoucherByCode(req.VoucherCode, "gift")
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fetchErr = errors.New("voucherCode not found")
			} else {
				fetchErr = err
				log.Error("Failed to fetch voucher", zap.Error(err))
			}
		}
		voucher = voucherData
	}()

	wg.Wait()

	if fetchErr != nil {
		return utils.ResponseErrors(ctx, fiber.StatusInternalServerError, fetchErr.Error())
	}

	result, err := voucherUsed.GetVoucherUsedByUserID(voucher.ID, user.ID)
	//if err != nil {
	//	log.Error("Failed to get voucher used by user", zap.Error(err))
	//	return utils.ResponseErrors(ctx, fiber.StatusInternalServerError, "operation failed(20156)")
	//}
	if result.RowsAffected > 0 {
		return utils.ResponseErrors(ctx, fiber.StatusConflict, "You have already used this voucher code")
	}

	if voucher.UsedCount == voucher.Reusability {
		return utils.ResponseErrors(ctx, fiber.StatusConflict, "voucherCode has been exceeded")
	}

	db = db.Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
			log.Error("Transaction failed", zap.Any("recover", r))
		}
	}()

	if err = transaction.Insert(user.ID, voucher.GiftAmount, "increase", "افزایش موجودی کیف پول از طریق هدیه"); err != nil {
		db.Rollback()
		return utils.ResponseErrors(ctx, fiber.StatusInternalServerError, "operation failed(20152)")
	}

	if err = userModel.UpdateUserBalance(user.ID, voucher.GiftAmount); err != nil {
		db.Rollback()
		return utils.ResponseErrors(ctx, fiber.StatusInternalServerError, "operation failed(20153)")
	}

	if err = voucherUsed.Insert(user.ID, voucher.ID); err != nil {
		db.Rollback()
		return utils.ResponseErrors(ctx, fiber.StatusInternalServerError, "operation failed(20154)")
	}

	if err = voucherModel.IncreaseUsedCount(voucher.ID); err != nil {
		db.Rollback()
		return utils.ResponseErrors(ctx, fiber.StatusInternalServerError, "operation failed(20155)")
	}

	db.Commit()

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"message": strings.NewReplacer("{GiftAmount}", strconv.FormatInt(voucher.GiftAmount, 10)).Replace("The amount of {GiftAmount} was added to your wallet"),
		},
	})
}
