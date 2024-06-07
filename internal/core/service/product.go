package service

import (
	"demo-hex-go/internal/core/domain/entity"
	"demo-hex-go/internal/core/port"

	"github.com/google/uuid"
)

type ProductService struct {
	repo port.ProductRepository
}
type ProductMongoService struct {
	repo port.ProductRepository
}

func NewProductService(repo port.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

/* Postgres Service */
func NewProductMongoService(repo port.ProductRepository) *ProductMongoService {
	return &ProductMongoService{
		repo: repo,
	}
}
func (productService *ProductService) SaveProduct(product entity.Product) error {
	product.Id = uuid.New().String()
	return productService.repo.SaveProduct(product)
}
func (productService *ProductService) ReadProducts() ([]*entity.Product, error) {
	return productService.repo.ReadProducts()
}

func (productService *ProductService) ReadProduct(id string) (*entity.Product, error) {
	return productService.repo.ReadProduct(id)
}
func (productService *ProductService) UpdateProduct(product entity.Product) error {
	return productService.repo.UpdateProduct(product)
}

func (productService *ProductService) DeleteProduct(id string) error {
	return productService.repo.DeleteProduct(id)
}

/**************************************************************************/
/* Mongo Service */
func (mongoService *ProductMongoService) SaveProduct(prouct entity.Product) error {
	prouct.Id = uuid.New().String()
	return mongoService.repo.SaveProduct(prouct)
}

func (mongoService *ProductMongoService) ReadProducts() ([]*entity.Product, error) {
	return mongoService.repo.ReadProducts()
}

func (mongoService *ProductMongoService) ReadProduct(id string) (*entity.Product, error) {
	return mongoService.repo.ReadProduct(id)
}
func (mongoService *ProductMongoService) UpdateProduct(product entity.Product) error {
	return mongoService.repo.UpdateProduct(product)
}

func (mongoService *ProductMongoService) DeleteProduct(id string) error {
	return mongoService.repo.DeleteProduct(id)
}
