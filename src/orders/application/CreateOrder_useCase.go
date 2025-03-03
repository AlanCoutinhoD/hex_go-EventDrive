package application

import (
	"demo/src/orders/domain/entities"
	"demo/src/orders/domain/repositories"
)

type CreateOrder struct {
	repository repositories.OrderRepository
}

func NewCreateOrder(repository repositories.OrderRepository) *CreateOrder {
	return &CreateOrder{
		repository: repository,
	}
}

func (co *CreateOrder) Execute(idProduct int, quantity int) error {
	order := entities.NewOrder(idProduct, quantity)
	return co.repository.Create(*order)
}
