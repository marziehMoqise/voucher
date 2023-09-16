package voucherUsedSchema

type ListRequest struct {
	Mobile      string `json:"mobile" validate:""`
	VoucherCode string `json:"voucherCode"`
}
