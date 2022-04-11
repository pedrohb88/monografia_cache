package transport

import (
	"context"
	"monografia/model"
	pb "monografia/transport/proto"
)

func (s *server) GetOrderByID(ctx context.Context, in *pb.ByIDRequest) (*pb.Order, error) {
	return s.entity.NewOrderByID(int(in.Id))
}

func (s *server) GetOrdersByUserID(ctx context.Context, in *pb.ByIDRequest) (*pb.Orders, error) {

	ordersModels, err := s.service.Orders.GetByUserID(int(in.Id))
	if err != nil {
		return nil, err
	}

	orders := make([]*pb.Order, len(ordersModels))
	for i, m := range ordersModels {

		order, err := s.entity.NewOrderByID(m.ID)
		if err != nil {
			return nil, err
		}

		orders[i] = order
	}

	return &pb.Orders{Orders: orders}, nil
}

func (s *server) CreateOrder(ctx context.Context, in *pb.Order) (*pb.Order, error) {

	order := model.Order{
		UserID:        int(in.UserId),
		ItemsQuantity: int(in.ItemsQuantity),
		Price:         float64(in.Price),
	}
	err := s.service.Orders.Create(&order)
	if err != nil {
		return nil, err
	}

	return s.entity.NewBasicOrder(&order), nil
}

func (s *server) AddItem(ctx context.Context, in *pb.Item) (*pb.Order, error) {

	err := s.service.Orders.AddItem(&model.Item{
		OrderID:   int(in.OrderId),
		ProductID: int(in.ProductId),
		Quantity:  int(in.Quantity),
		Price:     float64(in.Price),
	})
	if err != nil {
		return nil, err
	}

	return s.entity.NewOrderByID(int(in.OrderId))
}

func (s *server) RemoveItem(ctx context.Context, in *pb.ByIDRequest) (*pb.Order, error) {

	item, err := s.service.Items.GetByID(int(in.Id))
	if err != nil {
		return nil, err
	}

	err = s.service.Orders.RemoveItem(item.ID)
	if err != nil {
		return nil, err
	}

	return s.entity.NewOrderByID(item.OrderID)
}
