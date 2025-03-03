package application

import (
	"demo/src/products/domain/repositories"
	"demo/src/products/domain/entities"
)

type CreateProduct struct {
	repository domain.IProductRepository
}

func NewCreateProduct(repository domain.IProductRepository) *CreateProduct {
	return &CreateProduct{repository: repository}
}

func (cp *CreateProduct) Execute(product entities.Product) error {
	return cp.repository.Create(product)
}
