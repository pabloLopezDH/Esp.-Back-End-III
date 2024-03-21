package product

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/store"
)

type Repository interface {
	// GetByID busca un producto por su id
	GetByID(id int) (domain.Product, error)
	// Create agrega un nuevo producto
	Create(p domain.Product) (domain.Product, error)
	// Update actualiza un producto
	Update(id int, p domain.Product) (domain.Product, error)
	// Delete elimina un producto
	Delete(id int) error

	GetAll() ([]domain.Product, error)

	Search(params url.Values) ([]domain.Product, error)
}

type repository struct {
	storage store.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Product, error) {
	product, err := r.storage.Read(id)
	if err != nil {
		return domain.Product{}, errors.New("product not found")
	}
	return product, nil

}


func (r *repository) GetAll() ([]domain.Product, error) {
	products, err := r.storage.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil

}
func (r *repository) Search(params url.Values) ([]domain.Product, error) {

	var prod domain.ProductSearch
	if id := params.Get("id"); id != "" {
		parsedID, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		prod.Id = &parsedID
	}

	if name := params.Get("name"); name != "" {
		prod.Name = &name
	}
	
	if codeValue := params.Get("code_value"); codeValue != "" {
		prod.CodeValue = &codeValue
	}

	if isPublished := params.Get("is_published"); isPublished != "" {
		isPublishParsed, err := strconv.ParseBool(isPublished)
		if err != nil {
			return nil, err
		}
		prod.IsPublished = &isPublishParsed
	}

	products, err := r.storage.Search(prod)
	if err != nil {
		return nil, err
	}
	return products, nil

}

func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.storage.Exists(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	err := r.storage.Create(p)
	if err != nil {
		return domain.Product{}, errors.New("error creating product")
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, p domain.Product) (domain.Product, error) {
	if !r.storage.Exists(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	err := r.storage.Update(p)
	if err != nil {
		return domain.Product{}, errors.New("error updating product")
	}
	return p, nil
}
