package application

import (
	"demo/src/products/domain/repositories"
	"demo/src/products/domain/entities"
	"strconv"
)

type ListProductForId struct {
	repository domain.IProductRepository
}

func NewListProductForId(repository domain.IProductRepository) *ListProductForId {
	return &ListProductForId{repository: repository}
}

func (lp *ListProductForId) Execute(id string) (*entities.Product, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return lp.repository.GetByID(intID)
}
