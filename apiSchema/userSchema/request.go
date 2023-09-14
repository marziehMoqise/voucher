package userSchema

type GiftRequest struct {
	Mobile      string `json:"mobile" validate:"required,max=20"`
	VoucherCode string `json:"voucherCode" validate:"required"`
}
