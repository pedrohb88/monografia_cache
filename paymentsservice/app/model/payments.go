package model

type Payment struct {
	ID        int     `json:"id"`
	Amount    float64 `json:"amount"`
	InvoiceID *int    `json:"invoice_id"`
}
