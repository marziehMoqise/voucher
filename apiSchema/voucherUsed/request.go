package voucherUsedSchema

type ListRequest struct {
	Mobile      string `json:"mobile" validate:"max=20"`
	VoucherCode string `json:"voucherCode"`
}
