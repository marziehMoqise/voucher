package transactionSchema

type ListRequest struct {
	Mobile      string `json:"mobile" validate:"required,max=20"`
}
