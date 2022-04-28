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
	"google.golang.org/grpc/metadata"
)

type Payments interface {
	Create(ctx context.Context, amount float64) (int, error)
	GetByID(ctx context.Context, paymentID int) (*model.Payment, error)
}

type payments struct {
	url string
}

func New() Payments {
	return &payments{
		url: os.Getenv("PAYMENTS_URL"),
	}
}

func (p *payments) Create(ctx context.Context, amount float64) (int, error) {

	testID, reqID, err := getHeaders(ctx)
	if err != nil {
		return 0, err
	}

	conn, err := grpc.Dial(p.url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	c := pb.NewRouterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	md := metadata.New(map[string]string{
		"x-test": testID,
		"x-req":  reqID,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	payment, err := c.CreatePayment(ctx, &pb.Payment{
		Amount: float32(amount),
	})

	if err != nil {
		return 0, err
	}

	return int(payment.Id), err
}

func (p *payments) GetByID(ctx context.Context, paymentID int) (*model.Payment, error) {

	testID, reqID, err := getHeaders(ctx)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(p.url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewRouterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	md := metadata.New(map[string]string{
		"x-test": testID,
		"x-req":  reqID,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

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

func getHeaders(ctx context.Context) (string, string, error) {

	if os.Getenv("ENV") != "production" {
		return "", "", nil
	}

	testID := ctx.Value("x-test").(string)
	if testID == "" {
		return "", "", fmt.Errorf("missing x-test header")
	}

	reqID := ctx.Value("x-req").(string)
	if reqID == "" {
		return "", "", fmt.Errorf("missing x-req header")
	}
	return testID, reqID, nil
}
