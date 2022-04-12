package payments

import (
	"context"
	"fmt"
	"os"
	"time"

	"monografia/model"
	pb "monografia/transport/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Payments interface {
	Create(amount float64) (int, error)
	GetByID(paymentID int) (*model.Payment, error)
}

type payments struct {
	url string
}

func New() Payments {
	return &payments{
		url: os.Getenv("PAYMENTS_URL"),
	}
}

func (p *payments) Create(amount float64) (int, error) {
	fmt.Println("opaaaa")
	conn, err := grpc.Dial(p.url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	c := pb.NewRouterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	payment, err := c.CreatePayment(ctx, &pb.Payment{
		Amount: float32(amount),
	})

	if err != nil {
		return 0, err
	}

	return int(payment.Id), err
}

func (p *payments) GetByID(paymentID int) (*model.Payment, error) {

	conn, err := grpc.Dial(p.url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewRouterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	payment, err := c.GetPaymentByID(ctx, &pb.ByIDRequest{
		Id: int64(paymentID),
	})
	if err != nil {
		return nil, err
	}

	var invoice *model.Invoice
	var invoiceID *int
	if payment.Invoice != nil {
		invoice = &model.Invoice{
			ID:   int(payment.Invoice.Id),
			Code: payment.Invoice.Code,
			Link: payment.Invoice.Link,
		}
		invoiceID = &invoice.ID
	}

	return &model.Payment{
		ID:        int(payment.Id),
		Amount:    float64(payment.Amount),
		InvoiceID: invoiceID,
		Invoice:   invoice,
	}, nil
}
