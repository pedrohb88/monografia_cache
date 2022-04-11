package entity

import (
	"monografia/model"
	pb "monografia/transport/proto"
)

func (e *Entity) NewBasicInvoice(m *model.Invoice) *pb.Invoice {
	return &pb.Invoice{
		Id:   int64(m.ID),
		Code: m.Code,
		Link: m.Link,
	}
}

func (e *Entity) NewInvoiceByID(invoiceID int) (*pb.Invoice, error) {
	invoice, err := e.service.Invoices.GetByID(invoiceID)
	if err != nil {
		return nil, err
	}

	return e.NewBasicInvoice(invoice), nil
}
