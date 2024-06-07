package port

import "demo-hex-go/internal/core/domain/entity"

type ProductRepository interface {
	SaveProduct(product entity.Product) error
	ReadProducts() ([]*entity.Product, error)
	ReadProduct(id string) (*entity.Product, error)
	UpdateProduct(product entity.Product) error
	DeleteProduct(id string) error
}

type ProductService interface {
	SaveProduct(product entity.Product) error
	ReadProducts() ([]*entity.Product, error)
	ReadProduct(id string) (*entity.Product, error)
	UpdateProduct(product entity.Product) error
	DeleteProduct(id string) error
}
