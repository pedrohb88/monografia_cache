package entity

import (
	"context"
	"errors"
	libErrors "monografia/lib/errors"
	"monografia/model"
	pb "monografia/transport/proto"
)

func (e *Entity) NewOrderByID(ctx context.Context, orderID int) (*pb.Order, error) {
	orderModel, err := e.service.Orders.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	itemsModels, err := e.service.Items.GetByOrder(orderID)
	if err != nil && !errors.Is(err, libErrors.ErrNotFound) {
		return nil, err
	}

	items := make([]*pb.Item, len(itemsModels))
	for i, m := range itemsModels {
		item, err := e.NewItemByID(m.ID)
		if err != nil {
			return nil, err
		}
		items[i] = item
	}

	order := e.NewBasicOrder(orderModel)
	order.Items = items

	if order.PaymentId != 0 {
		paymentModel, err := e.service.Orders.GetPayment(ctx, int(order.PaymentId))
		if err != nil {
			return nil, err
		}

		var invoice *pb.Invoice
		var invoiceID int64
		if paymentModel.Invoice != nil {
			invoice = &pb.Invoice{
				Id:   int64(paymentModel.Invoice.ID),
				Code: paymentModel.Invoice.Code,
				Link: paymentModel.Invoice.Link,
			}
			invoiceID = invoice.Id
		}

		order.Payment = &pb.Payment{
			Id:        int64(paymentModel.ID),
			Amount:    float32(paymentModel.Amount),
			InvoiceId: invoiceID,
			Invoice:   invoice,
		}
	}

	return order, nil
}

func (e *Entity) NewBasicOrder(m *model.Order) *pb.Order {
	var paymentID int
	if m.PaymentID != nil {
		paymentID = *m.PaymentID
	}

	return &pb.Order{
		Id:            int64(m.ID),
		UserId:        int64(m.UserID),
		ItemsQuantity: int64(m.ItemsQuantity),
		Price:         float32(m.Price),
		PaymentId:     int64(paymentID),
	}
}
