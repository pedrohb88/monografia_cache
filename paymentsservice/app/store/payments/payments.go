package payments

import (
	"fmt"
	db "monografia/lib/database"
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	queryPaymentsByIDs = `
	SELECT 
		p.id AS ID,
		p.amount AS Amount,
		p.invoice_id AS InvoiceID
	FROM payments p
	WHERE p.id IN(%s)
	`

	execInsertPayment = `
	INSERT INTO payments(amount, invoice_id) 
	VALUES(?, ?)
	`
)

type Payments interface {
	GetByIDs(paymentIDs ...int) ([]*model.Payment, error)
	Create(payment *model.Payment) error
}

type payments struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Payments {
	return &payments{db: db}
}

func (i *payments) GetByIDs(paymentIDs ...int) ([]*model.Payment, error) {
	if len(paymentIDs) == 0 {
		return nil, nil
	}

	var payments []*model.Payment

	query := fmt.Sprintf(queryPaymentsByIDs, db.RepeatIntArgs(paymentIDs...))

	iIDS := make([]interface{}, len(paymentIDs))
	for i, id := range paymentIDs {
		iIDS[i] = id
	}

	_, err := i.db.Select(&payments, query, iIDS...)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (i *payments) Create(payment *model.Payment) error {
	res, err := i.db.Exec(execInsertPayment,
		payment.Amount,
		payment.InvoiceID,
	)
	if err != nil {
		return err
	}

	lastID, _ := res.LastInsertId()
	payment.ID = int(lastID)
	return nil
}
