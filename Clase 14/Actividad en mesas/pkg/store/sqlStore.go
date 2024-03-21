package store

import (
	"database/sql"
	"fmt"
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
	var product domain.Product
	query := "SELECT * FROM products WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *sqlStore) GetAll() ([]domain.Product, error) {
	listReturn := []domain.Product{}
	query := "SELECT * FROM products;"

	rows, err := s.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(product)
		listReturn = append(listReturn, product)
	}

	return listReturn, nil
}

func (s *sqlStore) Search(productSearch domain.ProductSearch) ([]domain.Product, error) {
	listReturn := []domain.Product{}
	q := "SELECT * FROM products WHERE "
	args := []interface{}{}

	// Add conditional query/args
	if productSearch.Id != nil {
		q = fmt.Sprintf("%s id=?", q)
		args = append(args, productSearch.Id)
	}

	if productSearch.Name != nil {
		if len(args) > 0 {
			q = fmt.Sprintf("%s AND ", q)
		} 
		q = fmt.Sprintf("%s name=?", q)
		args = append(args, productSearch.Name)
	}

	if productSearch.CodeValue != nil {
		if len(args) > 0 {
			q = fmt.Sprintf("%s AND ", q)
		} 
		q = fmt.Sprintf(" %s code_value=?", q)
		args = append(args, productSearch.CodeValue)
	}

	if productSearch.IsPublished != nil {
		if len(args) > 0 {
			q = fmt.Sprintf("%s AND ", q)
		} else {
			q = fmt.Sprintf(" %s is_published=?", q)
		}
		args = append(args, productSearch.IsPublished)
	}

	// Perform the query
	rows, err := s.db.Query(q, args...)
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(product)
		listReturn = append(listReturn, product)
	}
	return listReturn, nil
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
