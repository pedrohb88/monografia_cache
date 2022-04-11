package service

import (
	"monografia/lib/errors"
	"monografia/model"
	"monografia/store/invoices"
)

type invoicesService struct {
	invoicesStore invoices.Invoices
}

func (p *invoicesService) GetByID(invoiceID int) (*model.Invoice, error) {

	invoices, err := p.invoicesStore.GetByIDs(invoiceID)
	if err != nil {
		return nil, err
	}
	if len(invoices) == 0 {
		return nil, errors.ErrNotFound
	}

	return invoices[0], nil
}

func (p *invoicesService) Create(invoice *model.Invoice) error {
	return p.invoicesStore.Create(invoice)
}
