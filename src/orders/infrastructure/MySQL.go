package infrastructure

import (
	"database/sql"
	"demo/src/core"
	"demo/src/orders/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (m *MySQL) GetAll() []entities.Order {
	query := "SELECT id, idProduct, idClient, quantity FROM `order`"
	rows, err := m.conn.FetchRows(query)
	if err != nil {
		log.Fatalf("Error al obtener órdenes: %v", err)
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var order entities.Order
		if err := rows.Scan(&order.ID, &order.IdProduct, &order.IdClient, &order.Quantity); err != nil {
			log.Printf("Error al escanear la orden: %v", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error al iterar sobre las filas: %w", err)
	}

	return orders
}

func (m *MySQL) GetByID(id int) (*entities.Order, error) {
	query := "SELECT id, idProduct, idClient, quantity FROM `order` WHERE id = ?"
	log.Printf("Ejecutando consulta: %s con ID: %d", query, id)
	rows, err := m.conn.FetchRows(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var order entities.Order
	if rows.Next() {
		err = rows.Scan(&order.ID, &order.IdProduct, &order.IdClient, &order.Quantity)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("Orden no encontrada")
			}
			return nil, err
		}
		return &order, nil
	}

	return nil, fmt.Errorf("Orden no encontrada")
}

func (m *MySQL) Create(order entities.Order) error {
	query := "INSERT INTO `order` (idProduct, idClient, quantity) VALUES (?, ?, ?)"
	_, err := m.conn.ExecutePreparedQuery(query, order.IdProduct, order.IdClient, order.Quantity)
	if err != nil {
		log.Printf("Error al crear la orden: %v", err)
		return err
	}
	return nil
}

func (m *MySQL) Update(order entities.Order) error {
	query := "UPDATE `order` SET idProduct = ?, idClient = ?, quantity = ? WHERE id = ?"
	_, err := m.conn.ExecutePreparedQuery(query, order.IdProduct, order.IdClient, order.Quantity, order.ID)
	if err != nil {
		log.Printf("Error al actualizar la orden: %v", err)
		return err
	}
	log.Printf("Orden actualizada correctamente: %d", order.ID)
	return nil
}

func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM `order` WHERE id = ?"
	_, err := m.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("Error al eliminar la orden: %v", err)
		return err
	}
	return nil
}
