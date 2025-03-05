package application

import (
	"demo/src/orders/application/ports"
	"demo/src/orders/domain/entities"
	"log"
)

type CreateOrder struct {
	repository    ports.OrderRepository
	messageBroker ports.MessageBroker
}

func NewCreateOrder(repository ports.OrderRepository, messageBroker ports.MessageBroker) *CreateOrder {
	return &CreateOrder{
		repository:    repository,
		messageBroker: messageBroker,
	}
}

func (co *CreateOrder) Execute(idProduct int, idClient string, quantity int) error {
	order := entities.NewOrder(idProduct, idClient, quantity)

	err := co.repository.Create(*order)
	if err != nil {
		return err
	}

	if err := co.messageBroker.PublishOrder(*order); err != nil {
		log.Printf("Error publicando orden en RabbitMQ: %v", err)
	}

	return nil
}
