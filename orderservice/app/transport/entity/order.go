package entity

import (
	"errors"
	"fmt"
	libErrors "monografia/lib/errors"
	"monografia/model"
	pb "monografia/transport/proto"
)

func (e *Entity) NewOrderByID(orderID int) (*pb.Order, error) {
	orderModel, err := e.service.Orders.GetByID(orderID)
	if err != nil {
		fmt.Println("um")
		return nil, err
	}

	itemsModels, err := e.service.Items.GetByOrder(orderID)
	if err != nil && !errors.Is(err, libErrors.ErrNotFound) {
		fmt.Println("dois")

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
