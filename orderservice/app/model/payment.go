package model

type Invoice struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Link string `json:"link"`
}

type Payment struct {
	ID        int      `json:"id"`
	Amount    float64  `json:"amount"`
	InvoiceID *int     `json:"invoice_id"`
	Invoice   *Invoice `json:"invoice"`
}
