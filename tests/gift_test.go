package tests

import (
	"apiGolang/apiSchema/userSchema"
	"testing"
)

func TestGift(t *testing.T) {
	reqGift := userSchema.GiftRequest{
		Mobile:      "+989155099395",
		VoucherCode: "aaaa",
	}

	SendTestReq(reqGift, "/user/gift")
}
