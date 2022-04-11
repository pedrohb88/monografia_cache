package transport

import (
	"context"
	"monografia/model"
	pb "monografia/transport/proto"
)

func (s *server) GetPaymentByID(ctx context.Context, in *pb.ByIDRequest) (*pb.Payment, error) {
	return s.entity.NewPaymentByID(int(in.Id))
}

func (s *server) CreatePayment(ctx context.Context, in *pb.Payment) (*pb.Payment, error) {

	invoiceID := int(in.InvoiceId)

	payment := model.Payment{
		Amount:    float64(in.Amount),
		InvoiceID: &invoiceID,
	}
	err := s.service.Payments.Create(&payment)
	if err != nil {
		return nil, err
	}

	return s.entity.NewPaymentByID(payment.ID)
}
