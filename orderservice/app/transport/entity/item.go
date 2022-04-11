package entity

import (
	"monografia/model"
	pb "monografia/transport/proto"
)

func (e *Entity) NewBasicItem(m *model.Item) *pb.Item {
	return &pb.Item{
		Id:        int64(m.ID),
		OrderId:   int64(m.OrderID),
		ProductId: int64(m.ProductID),
		Quantity:  int64(m.Quantity),
		Price:     float32(m.Price),
	}
}

func (e *Entity) NewItemByID(itemID int) (*pb.Item, error) {

	itemModel, err := e.service.Items.GetByID(itemID)
	if err != nil {
		return nil, err
	}

	productModel, err := e.service.Products.GetByID(itemModel.ProductID)
	if err != nil {
		return nil, err
	}

	item := e.NewBasicItem(itemModel)
	item.Product = e.NewBasicProduct(productModel)
	return item, nil
}
