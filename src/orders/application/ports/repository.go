package ports

import "demo/src/orders/domain/entities"

type OrderRepository interface {
	Create(order entities.Order) error
	GetAll() []entities.Order
	GetByID(id int) (*entities.Order, error)
	Update(order entities.Order) error
}
