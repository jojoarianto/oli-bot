package model

// Payment data entity
type Payment struct {
	PaymentProof          string `json:"payment_proof"`
	PaymentProofThumbnail string `json:"payment_proof_thumbnail"`
	AccountName           string `json:"account_name"`
	PaymentStatus         int    `json:"payment_status" bson:"payment_status"`
	BankName              string `json:"bank_name"`
	TransferAt            string `json:"transfered_at"`
	Total                 string `json:"total"`
}
