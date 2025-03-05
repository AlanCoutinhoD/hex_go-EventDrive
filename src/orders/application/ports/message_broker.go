package ports

import "demo/src/orders/domain/entities"

type MessageBroker interface {
	PublishOrder(order entities.Order) error
}
