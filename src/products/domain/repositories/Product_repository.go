package domain

import "demo/src/products/domain/entities"

type IProduct interface {
	Save()
	GetAll() []entities.Product
	GetByID(id string) entities.Product
	Create(product entities.Product) error
	Delete(product entities.Product) error
}

type IProductRepository interface {
	GetAll() []entities.Product
	GetByID(id int) (*entities.Product, error)
	Create(product entities.Product) error
	Delete(id int) error
	Update(product entities.Product) error
}
