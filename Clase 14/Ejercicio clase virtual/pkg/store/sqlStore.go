package store

import (
	"database/sql"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) Read(id int) (domain.Product, error) {
	var product domain.Product
	query := "SELECT * FROM products WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *sqlStore) Create(product domain.Product) error {
	query := "INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Update(product domain.Product) error {
	query := "UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price, product.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Delete(id int) error {
	query := "DELETE FROM products WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Exists(codeValue string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM products WHERE code_value = ?;"
	row := s.db.QueryRow(query, codeValue)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
