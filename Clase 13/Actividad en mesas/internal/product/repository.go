package product

import (
	"errors"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
	"github.com/bootcamp-go/Consignas-Go-Web.git/pkg/store"
)

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) []domain.Product
	ConsumerPrice(list []int) ([]domain.Product, float64, error)
	Create(p domain.Product) (domain.Product, error)
	Update(id int, p domain.Product) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	storage store.Store
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.Store) Repository {
	return &repository{storage}
}

// GetAll devuelve todos los productos
func (r *repository) GetAll() []domain.Product {
	products, err := r.storage.GetAll()
	if err != nil {
		return []domain.Product{}
	}
	return products
}

// GetByID busca un producto por su id
func (r *repository) GetByID(id int) (domain.Product, error) {
	product, err := r.storage.GetOne(id)
	if err != nil {
		return domain.Product{}, errors.New("product not found")
	}
	return product, nil

}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (r *repository) SearchPriceGt(price float64) []domain.Product {
	var products []domain.Product
	list, err := r.storage.GetAll()
	if err != nil {
		return products
	}
	for _, product := range list {
		if product.Price > price {
			products = append(products, product)
		}
	}
	return products
}

// Create agrega un nuevo producto
func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	err := r.storage.AddOne(p)
	if err != nil {
		return domain.Product{}, errors.New("error creating product")
	}
	return p, nil
}

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *repository) validateCodeValue(codeValue string) bool {
	list, err := r.storage.GetAll()
	if err != nil {
		return false
	}
	for _, product := range list {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}

// Delete elimina un producto
func (r *repository) Delete(id int) error {
	err := r.storage.DeleteOne(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un producto
func (r *repository) Update(id int, p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	err := r.storage.UpdateOne(p)
	if err != nil {
		return domain.Product{}, errors.New("error updating product")
	}
	return p, nil
}

// ConsumerPrice devuelve una lista de productos y el precio total
func (r *repository) ConsumerPrice(list []int) ([]domain.Product, float64, error) {
	var products []domain.Product
	var total float64
	orders := make(map[int]int)
	for _, id := range list {
		if _, ok := orders[id]; ok {
			orders[id]++
		} else {
			orders[id] = 1
		}
	}
	for id, q := range orders {
		product, err := r.storage.GetOne(id)
		if err != nil {
			return products, total, errors.New("product not found")
		}
		if q > product.Quantity {
			return products, total, errors.New("not enough quantity")
		}
		for i := 0; i < q; i++ {
			products = append(products, product)
			total += product.Price
		}
	}
	return products, total, nil
}
