package application

import "demo/src/products/domain/entities"
import "demo/src/products/domain/repositories"


type ListProduct struct {
	repository domain.IProductRepository
}

func NewListProduct(repository domain.IProductRepository) *ListProduct {
	return &ListProduct{repository: repository}
}

func (lp *ListProduct) Execute() []entities.Product {
	return lp.repository.GetAll()
}