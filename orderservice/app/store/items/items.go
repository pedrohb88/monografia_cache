package items

import (
	"database/sql"
	"fmt"
	db "monografia/lib/database"
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	queryIDsByOrderID = `SELECT id FROM items WHERE order_id = ?`

	queryItemsByIDs = `
	SELECT 
		i.id AS ID,
		i.order_id AS OrderID,
		i.product_id AS ProductID,
		i.quantity AS Quantity,
		i.price AS Price
	FROM items i
	WHERE i.id IN(%s)
	`

	execInsertItem = `
	INSERT INTO items(order_id, product_id, quantity, price) 
	VALUES(?, ?, ?, ?)
	`

	execDeleteItem = `DELETE FROM items WHERE id = ?`
)

type Items interface {
	GetIDsByOrder(orderID int) ([]int, error)
	GetByIDs(itemIDs ...int) ([]*model.Item, error)
	Create(item *model.Item) error
	Delete(itemID int) error
}

type items struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Items {
	return &items{db: db}
}

func (i *items) GetIDsByOrder(orderID int) ([]int, error) {
	fmt.Println("pegando ids de items no repo")
	var ids []int
	_, err := i.db.Select(&ids, queryIDsByOrderID, orderID)
	return ids, err
}

func (i *items) GetByIDs(itemIDs ...int) ([]*model.Item, error) {
	fmt.Println("pegando os items no repo")
	if len(itemIDs) == 0 {
		return nil, nil
	}

	var items []*model.Item

	query := fmt.Sprintf(queryItemsByIDs, db.RepeatIntArgs(itemIDs...))

	iIDS := make([]interface{}, len(itemIDs))
	for i, id := range itemIDs {
		iIDS[i] = id
	}

	_, err := i.db.Select(&items, query, iIDS...)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *items) Create(item *model.Item) error {
	res, err := i.db.Exec(execInsertItem,
		item.OrderID,
		item.ProductID,
		item.Quantity,
		item.Price,
	)
	if err != nil {
		return err
	}

	lastID, _ := res.LastInsertId()
	item.ID = int(lastID)
	return nil
}

func (i *items) Delete(itemID int) error {
	res, err := i.db.Exec(execDeleteItem, itemID)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows > 0 {
		return nil
	}
	return sql.ErrNoRows
}
