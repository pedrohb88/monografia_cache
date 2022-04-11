package orders

import (
	"fmt"
	db "monografia/lib/database"
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	queryIDsByUser = `SELECT id FROM orders WHERE user_id = ?`

	queryByIDs = `
	SELECT 
		o.id AS ID, 
		o.user_id AS UserID, 
		o.items_quantity AS ItemsQuantity, 
		o.price AS Price,
		o.payment_id AS PaymentID
	FROM orders o
	WHERE o.id IN (%s)
	`

	execInsertOrder = `
	INSERT INTO orders(user_id, items_quantity, price)
	VALUES (?, ?, ?)
	`
)

type Orders interface {
	GetIDsByUser(userID int) ([]int, error)
	GetByIDs(orderIDs ...int) ([]*model.Order, error)
	Create(order *model.Order) error
}

type orders struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Orders {
	return &orders{db: db}
}

func (o *orders) GetIDsByUser(userID int) ([]int, error) {
	fmt.Println("pegando os ids de orders no repo")
	var ids []int
	_, err := o.db.Select(&ids, queryIDsByUser, userID)
	return ids, err
}

func (o *orders) GetByIDs(orderIDs ...int) ([]*model.Order, error) {
	fmt.Println("pegando as orders no repo")
	if len(orderIDs) == 0 {
		return nil, nil
	}

	var orders []*model.Order

	query := fmt.Sprintf(queryByIDs, db.RepeatIntArgs(orderIDs...))

	iIDS := make([]interface{}, len(orderIDs))
	for i, id := range orderIDs {
		iIDS[i] = id
	}

	_, err := o.db.Select(&orders, query, iIDS...)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orders) Create(order *model.Order) error {

	res, err := o.db.Exec(execInsertOrder,
		order.UserID,
		order.ItemsQuantity,
		order.Price,
	)
	if err != nil {
		return err
	}

	lastID, _ := res.LastInsertId()
	order.ID = int(lastID)
	return nil
}
