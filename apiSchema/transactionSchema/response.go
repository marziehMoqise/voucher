package transactionSchema

type ListResponse struct {
	ID          int64
	Amount      int64
	Type        string
	Time        int64
	Description string
}
