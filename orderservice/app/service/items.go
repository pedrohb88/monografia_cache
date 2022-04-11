package service

import (
	"monografia/lib/errors"
	"monografia/model"
	"monografia/store/items"
)

type itemsService struct {
	itemsStore items.Items
}

func (i *itemsService) GetByOrder(orderID int) ([]*model.Item, error) {
	ids, err := i.itemsStore.GetIDsByOrder(orderID)
	if err != nil {
		return nil, err
	}

	return i.itemsStore.GetByIDs(ids...)
}

func (i *itemsService) GetByID(itemID int) (*model.Item, error) {
	items, err := i.itemsStore.GetByIDs(itemID)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, errors.ErrNotFound
	}

	return items[0], nil
}
