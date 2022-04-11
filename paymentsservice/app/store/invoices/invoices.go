package invoices

import (
	"fmt"
	db "monografia/lib/database"
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	queryInvoicesByIDs = `
	SELECT 
		i.id AS ID,
		i.code AS Code,
		i.link AS Link
	FROM invoices i
	WHERE i.id IN(%s)
	`

	execInsert = `
	INSERT INTO invoices(code, link)
	VALUES (?, ?)
	`
)

type Invoices interface {
	GetByIDs(invoiceIDs ...int) ([]*model.Invoice, error)
	Create(invoice *model.Invoice) error
}

type invoices struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Invoices {
	return &invoices{db: db}
}

func (p *invoices) GetByIDs(invoiceIDs ...int) ([]*model.Invoice, error) {

	if len(invoiceIDs) == 0 {
		return nil, nil
	}

	var invoices []*model.Invoice

	query := fmt.Sprintf(queryInvoicesByIDs, db.RepeatIntArgs(invoiceIDs...))

	iIDS := make([]interface{}, len(invoiceIDs))
	for i, id := range invoiceIDs {
		iIDS[i] = id
	}

	_, err := p.db.Select(&invoices, query, iIDS...)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (p *invoices) Create(invoice *model.Invoice) error {

	res, err := p.db.Exec(execInsert,
		invoice.Code,
		invoice.Link,
	)
	if err != nil {
		return err
	}

	lastID, _ := res.LastInsertId()
	invoice.ID = int(lastID)
	return nil
}
