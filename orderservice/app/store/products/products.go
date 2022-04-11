package products

import (
	"database/sql"
	"fmt"
	db "monografia/lib/database"
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	queryAllIDs = `SELECT id FROM products`

	queryProductsByIDs = `
	SELECT 
		p.id AS ID,
		p.name AS Name,
		p.price AS Price
	FROM products p
	WHERE p.id IN(%s)
	`

	execInsert = `
	INSERT INTO products(name, price)
	VALUES (?, ?)
	`

	execDelete = `DELETE FROM products WHERE id = ?`
)

type Products interface {
	GetAllIDs() ([]int, error)
	GetByIDs(productIDs ...int) ([]*model.Product, error)
	Create(product *model.Product) error
	Delete(productID int) error
}

type products struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Products {
	return &products{db: db}
}

func (p *products) GetAllIDs() ([]int, error) {
	var ids []int
	_, err := p.db.Select(&ids, queryAllIDs)
	return ids, err
}

func (p *products) GetByIDs(productIDs ...int) ([]*model.Product, error) {

	if len(productIDs) == 0 {
		return nil, nil
	}

	var products []*model.Product

	query := fmt.Sprintf(queryProductsByIDs, db.RepeatIntArgs(productIDs...))

	iIDS := make([]interface{}, len(productIDs))
	for i, id := range productIDs {
		iIDS[i] = id
	}

	_, err := p.db.Select(&products, query, iIDS...)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *products) Create(product *model.Product) error {

	res, err := p.db.Exec(execInsert,
		product.Name,
		product.Price,
	)
	if err != nil {
		return err
	}

	lastID, _ := res.LastInsertId()
	product.ID = int(lastID)
	return nil
}

func (p *products) Delete(productID int) error {
	res, err := p.db.Exec(execDelete, productID)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows > 0 {
		return nil
	}
	return sql.ErrNoRows
}
