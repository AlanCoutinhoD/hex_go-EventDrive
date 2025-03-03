package application

import (
	"demo/src/products/domain/repositories"
)

type DeleteProduct struct {
	repository domain.IProductRepository
}

func NewDeleteProduct(repository domain.IProductRepository) *DeleteProduct {
	return &DeleteProduct{repository: repository}
}

func (dp *DeleteProduct) Execute(id int) error {
	return dp.repository.Delete(id)
}
