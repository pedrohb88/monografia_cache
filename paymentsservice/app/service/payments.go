package service

import (
	"monografia/lib/errors"
	"monografia/model"
	"monografia/store/payments"
)

type paymentsService struct {
	paymentsStore payments.Payments
}

func (i *paymentsService) GetByID(paymentID int) (*model.Payment, error) {
	payments, err := i.paymentsStore.GetByIDs(paymentID)
	if err != nil {
		return nil, err
	}
	if len(payments) == 0 {
		return nil, errors.ErrNotFound
	}

	return payments[0], nil
}

func (i *paymentsService) Create(payment *model.Payment) error {
	return i.paymentsStore.Create(payment)
}
