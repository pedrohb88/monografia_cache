package service

import (
	"monografia/lib/errors"
	"monografia/model"
	"monografia/store/items"
	"monografia/store/orders"
	"monografia/store/products"
)

type ordersService struct {
	ordersStore   orders.Orders
	itemsStore    items.Items
	productsStore products.Products
}

func (o *ordersService) GetByID(orderID int) (*model.Order, error) {
	orders, err := o.ordersStore.GetByIDs(orderID)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, errors.ErrNotFound
	}

	return orders[0], nil
}

func (o *ordersService) GetByUserID(userID int) ([]*model.Order, error) {
	ids, err := o.ordersStore.GetIDsByUser(userID)
	if err != nil {
		return nil, err
	}

	return o.ordersStore.GetByIDs(ids...)
}

func (o *ordersService) Create(order *model.Order) error {
	return o.ordersStore.Create(order)
}

func (o *ordersService) AddItem(item *model.Item) error {

	_, err := o.productsStore.GetByIDs(item.ProductID)
	if errors.IsNotFound(err) {
		return errors.ErrProductNotFound
	}
	if err != nil {
		return err
	}

	return o.itemsStore.Create(item)
}

func (o *ordersService) RemoveItem(itemID int) error {
	return o.itemsStore.Delete(itemID)
}
