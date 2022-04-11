package invoices

import (
	"monografia/lib/cache"
	"monografia/model"

	store "monografia/store/invoices"
)

type Invoices struct {
	cache *cache.Cache
	store store.Invoices
}

var digest = cache.CalculateDigest(model.Invoice{})

func byIDCacheKey(id int) string {
	return cache.Key("invoicebyid", digest, id)
}

func byIDsCacheKeys(ids ...int) []string {
	keys := make([]string, len(ids))

	for i, id := range ids {
		key := cache.Key("invoicebyid", digest, id)
		keys[i] = key
	}
	return keys
}

func New(cache *cache.Cache, store store.Invoices) store.Invoices {
	return &Invoices{cache: cache, store: store}
}

func (p *Invoices) GetByIDs(invoiceIDs ...int) ([]*model.Invoice, error) {

	invoicesMap := map[string]*model.Invoice{}
	keys := byIDsCacheKeys(invoiceIDs...)

	err := p.cache.GetMulti(keys, func(key string) interface{} {
		var emptyInvoice model.Invoice
		invoicesMap[key] = &emptyInvoice
		return invoicesMap[key]
	})
	if err != nil {
		return nil, err
	}

	var invoices []*model.Invoice
	var notFoundIDs []int
	for _, id := range invoiceIDs {
		key := byIDCacheKey(id)

		if v, ok := invoicesMap[key]; ok {
			invoices = append(invoices, v)
			continue
		}

		notFoundIDs = append(notFoundIDs, id)
	}

	var dbInvoices []*model.Invoice

	if len(notFoundIDs) > 0 {
		dbInvoices, err = p.store.GetByIDs(notFoundIDs...)
		if err != nil {
			return nil, err
		}

		for _, dbInvoice := range dbInvoices {
			err = p.cache.Set(byIDCacheKey(dbInvoice.ID), dbInvoice)
			if err != nil {
				return nil, err
			}
		}
	}

	return append(invoices, dbInvoices...), nil
}

func (p *Invoices) Create(invoice *model.Invoice) error {
	err := p.store.Create(invoice)
	if err != nil {
		return err
	}

	return p.cache.Delete(byIDCacheKey(invoice.ID))
}
