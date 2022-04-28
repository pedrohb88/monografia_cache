package orders

import (
	"monografia/lib/cache"
	"monografia/model"
	store "monografia/store/orders"
)

type Orders struct {
	cache *cache.Cache
	store store.Orders
}

var digest = cache.CalculateDigest(model.Order{})

func byIDCacheKey(id int) string {
	return cache.Key("orderbyid", digest, id)
}

func byIDsCacheKeys(ids ...int) []string {
	keys := make([]string, len(ids))

	for i, id := range ids {
		key := cache.Key("orderbyid", digest, id)
		keys[i] = key
	}
	return keys
}

func byUserCacheKey(userID int) string {
	return cache.Key("ordersbyuser", digest, userID)
}

func New(cache *cache.Cache, store store.Orders) store.Orders {
	return &Orders{cache: cache, store: store}
}

func (p *Orders) GetIDsByUser(userID int) ([]int, error) {
	var ids []int

	err := p.cache.GetSet(byUserCacheKey(userID), &ids, func() (interface{}, error) {
		return p.store.GetIDsByUser(userID)
	})

	return ids, err
}

func (p *Orders) GetByIDs(orderIDs ...int) ([]*model.Order, error) {
	ordersMap := map[string]*model.Order{}
	keys := byIDsCacheKeys(orderIDs...)

	err := p.cache.GetMulti(keys, func(key string) interface{} {
		var emptyOrder model.Order
		ordersMap[key] = &emptyOrder
		return ordersMap[key]
	})
	if err != nil {
		return nil, err
	}

	var orders []*model.Order
	var notFoundIDs []int
	for _, id := range orderIDs {
		key := byIDCacheKey(id)

		if v, ok := ordersMap[key]; ok {
			orders = append(orders, v)
			continue
		}

		notFoundIDs = append(notFoundIDs, id)
	}

	var dbOrders []*model.Order

	if len(notFoundIDs) > 0 {
		dbOrders, err = p.store.GetByIDs(notFoundIDs...)
		if err != nil {
			return nil, err
		}

		for _, dbOrder := range dbOrders {
			err = p.cache.Set(byIDCacheKey(dbOrder.ID), dbOrder)
			if err != nil {
				return nil, err
			}
		}
	}

	return append(orders, dbOrders...), nil
}

func (o *Orders) Create(order *model.Order) error {
	err := o.store.Create(order)
	if err != nil {
		return err
	}

	return o.cache.Delete(byUserCacheKey(order.UserID))
}

func (o *Orders) UpdatePaymentID(orderID, paymentID int) error {
	return o.store.UpdatePaymentID(orderID, paymentID)
}
