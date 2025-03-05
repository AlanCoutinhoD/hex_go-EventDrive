package rabbitmq

import (
	"encoding/json"
	"os"

	"demo/src/orders/domain/entities"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ() (*RabbitMQ, error) {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@localhost:5672/"
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	queueName := os.Getenv("RABBITMQ_QUEUE")
	if queueName == "" {
		queueName = "orders_queue"
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
	}, nil
}

func (r *RabbitMQ) PublishOrder(order entities.Order) error {
	queueName := os.Getenv("RABBITMQ_QUEUE")
	if queueName == "" {
		queueName = "orders_queue"
	}

	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = r.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	return err
}

func (r *RabbitMQ) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
