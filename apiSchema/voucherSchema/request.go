package voucherSchema

type SetRequest struct {
	VoucherCode string `json:"voucherCode" validate:"required,max=24"`
	GiftAmount  int64  `json:"giftAmount" validate:"required"`
	Reusability int64  `json:"reusability" validate:"required"`
}
