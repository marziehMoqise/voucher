package tests

import (
	voucherUsedSchema "apiGolang/apiSchema/voucherUsedSchema"
	"testing"
)

func TestVoucherUsedList(t *testing.T) {
	VoucherUsedListRequest := voucherUsedSchema.ListRequest{
		//Mobile:      "+989155099394",
		VoucherCode: "aaaa",
	}

	SendTestReq(VoucherUsedListRequest, "/vouchersUsed")
}
