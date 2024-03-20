package store

import (
	"database/sql"
	"log"

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
	return domain.Product{}, nil
}

// Create agrega un nuevo producto
func (s *sqlStore) Create(product domain.Product) error {
	query := "INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Update actualiza un producto
func (s *sqlStore) Update(product domain.Product) error {
	return nil
}

// Delete elimina un producto
func (s *sqlStore) Delete(id int) error {
	return nil
}

// Exists verifica si un producto existe
func (s *sqlStore) Exists(codeValue string) bool {
	return true
}
