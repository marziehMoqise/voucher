package tests

import (
	"apiGolang/apiSchema/transactionSchema"
	"testing"
)

func TestTransactionList(t *testing.T) {
	TransactionListRequest := transactionSchema.ListRequest{
		Mobile:      "+989155099394",
	}

	SendTestReq(TransactionListRequest, "/transactions")
}
