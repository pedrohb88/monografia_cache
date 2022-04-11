package service

import (
	"monografia/lib/errors"
	"monografia/model"
	"monografia/store/products"
)

type productsService struct {
	productsStore products.Products
}

func (p *productsService) GetAll() ([]*model.Product, error) {

	ids, err := p.productsStore.GetAllIDs()
	if err != nil {
		return nil, err
	}

	return p.productsStore.GetByIDs(ids...)
}

func (p *productsService) GetByID(productID int) (*model.Product, error) {

	products, err := p.productsStore.GetByIDs(productID)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.ErrNotFound
	}

	return products[0], nil
}

func (p *productsService) Create(product *model.Product) error {
	return p.productsStore.Create(product)
}

func (p *productsService) Delete(productID int) error {
	return p.productsStore.Delete(productID)
}
