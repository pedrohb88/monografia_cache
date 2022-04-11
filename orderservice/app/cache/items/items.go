package items

import (
	"fmt"
	"monografia/lib/cache"
	"monografia/model"
	store "monografia/store/items"
)

type Items struct {
	cache *cache.Cache
	store store.Items
}

var digest = cache.CalculateDigest(model.Item{})

func byOrderCacheKey(orderID int) string {
	return cache.Key("itemsbyorder", digest, orderID)
}

func byIDCacheKey(id int) string {
	return cache.Key("itembyid", digest, id)
}

func byIDsCacheKeys(ids ...int) []string {
	keys := make([]string, len(ids))

	for i, id := range ids {
		key := cache.Key("itembyid", digest, id)
		keys[i] = key
	}
	return keys
}

func New(cache *cache.Cache, store store.Items) store.Items {
	return &Items{cache: cache, store: store}
}

func (i *Items) GetIDsByOrder(orderID int) ([]int, error) {
	fmt.Println("pegando ids de items na cache")

	var ids []int

	err := i.cache.GetSet(byOrderCacheKey(orderID), &ids, func() (interface{}, error) {
		return i.store.GetIDsByOrder(orderID)
	})

	return ids, err
}

func (i *Items) GetByIDs(itemIDs ...int) ([]*model.Item, error) {
	fmt.Println("pegando items na cache")

	itemsMap := map[string]*model.Item{}
	keys := byIDsCacheKeys(itemIDs...)

	err := i.cache.GetMulti(keys, func(key string) interface{} {
		var emptyItem model.Item
		itemsMap[key] = &emptyItem
		return itemsMap[key]
	})
	if err != nil {
		return nil, err
	}

	var products []*model.Item
	var notFoundIDs []int
	for _, id := range itemIDs {
		key := byIDCacheKey(id)

		if v, ok := itemsMap[key]; ok {
			products = append(products, v)
			continue
		}

		notFoundIDs = append(notFoundIDs, id)
	}

	var dbItems []*model.Item

	if len(notFoundIDs) > 0 {
		dbItems, err = i.store.GetByIDs(notFoundIDs...)
		if err != nil {
			return nil, err
		}

		for _, dbItem := range dbItems {
			err = i.cache.Set(byIDCacheKey(dbItem.ID), dbItem)
			if err != nil {
				return nil, err
			}
		}
	}

	return append(products, dbItems...), nil
}

func (i *Items) Create(item *model.Item) error {
	err := i.store.Create(item)
	if err != nil {
		return err
	}

	return i.cache.Delete(byOrderCacheKey(item.OrderID))
}

func (i *Items) Delete(itemID int) error {
	items, err := i.store.GetByIDs(itemID)
	if err != nil {
		return err
	}

	err = i.store.Delete(itemID)
	if err != nil {
		return err
	}

	err = i.cache.Delete(byIDCacheKey(itemID))
	if err != nil {
		return err
	}
	return i.cache.Delete(byOrderCacheKey(items[0].OrderID))
}
