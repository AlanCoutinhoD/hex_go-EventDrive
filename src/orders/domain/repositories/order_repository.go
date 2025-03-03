package repositories

import "demo/src/orders/domain/entities"

type OrderRepository interface {
	GetAll() []entities.Order
	GetByID(id int) (*entities.Order, error)
	Create(order entities.Order) error
	Update(order entities.Order) error
	Delete(id int) error
}
