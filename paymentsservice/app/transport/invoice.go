package transport

import (
	"context"
	"monografia/model"
	pb "monografia/transport/proto"
)

func (s *server) GetInvoiceByID(ctx context.Context, in *pb.ByIDRequest) (*pb.Invoice, error) {
	return s.entity.NewInvoiceByID(int(in.Id))
}

func (s *server) CreateInvoice(ctx context.Context, in *pb.Invoice) (*pb.Invoice, error) {

	invoiceModel := model.Invoice{
		Code: in.Code,
		Link: in.Link,
	}

	err := s.service.Invoices.Create(&invoiceModel)
	if err != nil {
		return nil, err
	}

	return s.entity.NewBasicInvoice(&invoiceModel), nil
}
