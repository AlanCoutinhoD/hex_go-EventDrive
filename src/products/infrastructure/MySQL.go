package infrastructure

import (
	"demo/src/core"
	"demo/src/products/domain/entities"
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

func (mysql *MySQL) Save() {
	log.Println("Método Save() está deprecado. Por favor usar Create()")
}

func (m *MySQL) GetAll() []entities.Product {
	query := "SELECT id, name, price, description, image FROM product"
	rows, err := m.conn.FetchRows(query)
	if err != nil {
		log.Fatalf("Error al obtener productos: %v", err)
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Image); err != nil {
			log.Printf("Error al escanear el producto: %v", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error al iterar sobre las filas: %w", err)
	}

	return products
}

func (m *MySQL) GetByID(id int) (*entities.Product, error) {
	log.Printf("Buscando producto con ID: %d", id)
	query := "SELECT id, name, price, description, image FROM product WHERE id = ?"
	rows, err := m.conn.FetchRows(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var product entities.Product
	if rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Image); err != nil {
			return nil, err
		}
		return &product, nil
	}

	return nil, fmt.Errorf("Producto no encontrado")
}

func (m *MySQL) Create(product entities.Product) error {
	query := "INSERT INTO product (name, price, description, image) VALUES (?, ?, ?, ?)"

	result, err := m.conn.ExecutePreparedQuery(query, product.Name, product.Price, product.Description, product.Image)
	if err != nil {
		return fmt.Errorf("error al crear el producto: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Producto creado correctamente")
	}

	return nil
}

func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM product WHERE id = ?"
	_, err := m.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el producto: %v", err)
	}
	return nil
}

func (m *MySQL) Update(product entities.Product) error {
	query := "UPDATE product SET name = ?, price = ?, description = ?, image = ? WHERE id = ?"
	_, err := m.conn.ExecutePreparedQuery(query, product.Name, product.Price, product.Description, product.Image, product.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar el producto: %v", err)
	}
	return nil
}
