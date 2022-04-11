package transport

import (
	"context"
	"monografia/model"
	pb "monografia/transport/proto"
)

func (s *server) GetAllProducts(ctx context.Context, in *pb.EmptyRequest) (*pb.Products, error) {

	productsModels, err := s.service.Products.GetAll()
	if err != nil {
		return nil, err
	}

	products := make([]*pb.Product, len(productsModels))
	for i, v := range productsModels {
		products[i] = s.entity.NewBasicProduct(v)
	}

	return &pb.Products{Products: products}, nil
}

func (s *server) GetProductByID(ctx context.Context, in *pb.ByIDRequest) (*pb.Product, error) {
	return s.entity.NewProductByID(int(in.Id))
}

func (s *server) CreateProduct(ctx context.Context, in *pb.Product) (*pb.Product, error) {

	productModel := model.Product{
		Name:  in.Name,
		Price: float64(in.Price),
	}

	err := s.service.Products.Create(&productModel)
	if err != nil {
		return nil, err
	}

	return s.entity.NewBasicProduct(&productModel), nil
}

func (s *server) DeleteProduct(ctx context.Context, in *pb.ByIDRequest) (*pb.EmptyResponse, error) {
	return &pb.EmptyResponse{}, s.service.Products.Delete(int(in.Id))
}
