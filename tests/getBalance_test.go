package tests

import (
	"apiGolang/apiSchema/userSchema"
	"testing"
)

func TestGetBalance(t *testing.T) {

	BalanceRequestRequest := userSchema.BalanceRequest{
		Mobile:      "+989155099391",
	}

	SendTestReq(BalanceRequestRequest, "/user/getBalance")
}
