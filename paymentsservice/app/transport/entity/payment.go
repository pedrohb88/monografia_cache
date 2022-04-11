package entity

import (
	"monografia/model"
	pb "monografia/transport/proto"
)

func (e *Entity) NewBasicPayment(m *model.Payment) *pb.Payment {

	var invoiceID int
	if m.InvoiceID != nil {
		invoiceID = *m.InvoiceID
	}

	return &pb.Payment{
		Id:        int64(m.ID),
		Amount:    float32(m.Amount),
		InvoiceId: int64(invoiceID),
	}
}

func (e *Entity) NewPaymentByID(paymentID int) (*pb.Payment, error) {

	paymentModel, err := e.service.Payments.GetByID(paymentID)
	if err != nil {
		return nil, err
	}

	payment := e.NewBasicPayment(paymentModel)

	if paymentModel.InvoiceID != nil {
		invoiceModel, err := e.service.Invoices.GetByID(*paymentModel.InvoiceID)
		if err != nil {
			return nil, err
		}
		payment.Invoice = e.NewBasicInvoice(invoiceModel)
	}

	return payment, nil
}
