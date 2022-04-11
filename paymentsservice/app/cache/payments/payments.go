package payments

import (
	"monografia/lib/cache"
	"monografia/model"
	store "monografia/store/payments"
)

type Payments struct {
	cache *cache.Cache
	store store.Payments
}

var digest = cache.CalculateDigest(model.Payment{})

func byIDCacheKey(id int) string {
	return cache.Key("paymentbyid", digest, id)
}

func byIDsCacheKeys(ids ...int) []string {
	keys := make([]string, len(ids))

	for i, id := range ids {
		key := cache.Key("paymentbyid", digest, id)
		keys[i] = key
	}
	return keys
}

func New(cache *cache.Cache, store store.Payments) store.Payments {
	return &Payments{cache: cache, store: store}
}

func (i *Payments) GetByIDs(paymentIDs ...int) ([]*model.Payment, error) {
	paymentsMap := map[string]*model.Payment{}
	keys := byIDsCacheKeys(paymentIDs...)

	err := i.cache.GetMulti(keys, func(key string) interface{} {
		var emptyPayment model.Payment
		paymentsMap[key] = &emptyPayment
		return paymentsMap[key]
	})
	if err != nil {
		return nil, err
	}

	var invoices []*model.Payment
	var notFoundIDs []int
	for _, id := range paymentIDs {
		key := byIDCacheKey(id)

		if v, ok := paymentsMap[key]; ok {
			invoices = append(invoices, v)
			continue
		}

		notFoundIDs = append(notFoundIDs, id)
	}

	var dbPayments []*model.Payment

	if len(notFoundIDs) > 0 {
		dbPayments, err = i.store.GetByIDs(notFoundIDs...)
		if err != nil {
			return nil, err
		}

		for _, dbPayment := range dbPayments {
			err = i.cache.Set(byIDCacheKey(dbPayment.ID), dbPayment)
			if err != nil {
				return nil, err
			}
		}
	}

	return append(invoices, dbPayments...), nil
}

func (i *Payments) Create(payment *model.Payment) error {
	return i.store.Create(payment)
}
