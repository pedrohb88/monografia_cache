package products

import (
	"monografia/lib/cache"
	"monografia/model"

	store "monografia/store/products"
)

type Products struct {
	cache *cache.Cache
	store store.Products
}

var digest = cache.CalculateDigest(model.Product{})

func byIDCacheKey(id int) string {
	return cache.Key("productbyid", digest, id)
}

func byIDsCacheKeys(ids ...int) []string {
	keys := make([]string, len(ids))

	for i, id := range ids {
		key := cache.Key("productbyid", digest, id)
		keys[i] = key
	}
	return keys
}

func allCacheKey() string {
	return cache.Key("productsall", digest)
}

func New(cache *cache.Cache, store store.Products) store.Products {
	return &Products{cache: cache, store: store}
}

func (p *Products) GetAllIDs() ([]int, error) {

	var ids []int

	err := p.cache.GetSet(allCacheKey(), &ids, func() (interface{}, error) {
		return p.store.GetAllIDs()
	})

	return ids, err
}

func (p *Products) GetByIDs(productIDs ...int) ([]*model.Product, error) {

	productsMap := map[string]*model.Product{}
	keys := byIDsCacheKeys(productIDs...)

	err := p.cache.GetMulti(keys, func(key string) interface{} {
		var emptyProduct model.Product
		productsMap[key] = &emptyProduct
		return productsMap[key]
	})
	if err != nil {
		return nil, err
	}

	var products []*model.Product
	var notFoundIDs []int
	for _, id := range productIDs {
		key := byIDCacheKey(id)

		if v, ok := productsMap[key]; ok {
			products = append(products, v)
			continue
		}

		notFoundIDs = append(notFoundIDs, id)
	}

	var dbProducts []*model.Product

	if len(notFoundIDs) > 0 {
		dbProducts, err = p.store.GetByIDs(notFoundIDs...)
		if err != nil {
			return nil, err
		}

		for _, dbProduct := range dbProducts {
			err = p.cache.Set(byIDCacheKey(dbProduct.ID), dbProduct)
			if err != nil {
				return nil, err
			}
		}
	}

	return append(products, dbProducts...), nil
}

func (p *Products) Create(product *model.Product) error {
	err := p.store.Create(product)
	if err != nil {
		return err
	}

	err = p.cache.Delete(byIDCacheKey(product.ID))
	if err != nil {
		return err
	}
	return p.cache.Delete(allCacheKey())
}

func (p *Products) Delete(productID int) error {
	err := p.store.Delete(productID)
	if err != nil {
		return err
	}

	err = p.cache.Delete(byIDCacheKey(productID))
	if err != nil {
		return err
	}
	return p.cache.Delete(allCacheKey())
}
