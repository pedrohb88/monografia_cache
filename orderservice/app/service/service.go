package service

import (
	"monografia/store/items"
	"monografia/store/orders"
	"monografia/store/products"
)

type Service struct {
	Orders   ordersService
	Items    itemsService
	Products productsService
}

func New(
	ordersStore orders.Orders,
	productsStore products.Products,
	itemsStore items.Items,
) Service {

	ordersService := ordersService{
		ordersStore:   ordersStore,
		itemsStore:    itemsStore,
		productsStore: productsStore,
	}

	itemsService := itemsService{
		itemsStore: itemsStore,
	}

	productsService := productsService{
		productsStore: productsStore,
	}

	return Service{
		Orders:   ordersService,
		Items:    itemsService,
		Products: productsService,
	}
}
