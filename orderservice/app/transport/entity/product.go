package entity

import (
	"monografia/model"
	pb "monografia/transport/proto"
)

func (e *Entity) NewBasicProduct(m *model.Product) *pb.Product {
	return &pb.Product{
		Id:    int64(m.ID),
		Name:  m.Name,
		Price: float32(m.Price),
	}
}

func (e *Entity) NewProductByID(productID int) (*pb.Product, error) {
	product, err := e.service.Products.GetByID(productID)
	if err != nil {
		return nil, err
	}

	return e.NewBasicProduct(product), nil
}
