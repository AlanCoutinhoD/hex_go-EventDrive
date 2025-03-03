package infrastructure

import (
	"database/sql"
	"demo/src/core"
	"demo/src/employees/domain/entities"
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

func (m *MySQL) Create(employee entities.Employee) {
	query := "INSERT INTO employees (name) VALUES (?)"
	fmt.Println(employee)
	_, err := m.conn.ExecutePreparedQuery(query, employee.Name)
	if err != nil {
		log.Fatalf("Error al crear el empleado: %v", err)
	}
}

func (m *MySQL) GetAll() []entities.Employee {
	query := "SELECT * FROM employees"
	rows, err := m.conn.FetchRows(query)
	if err != nil {
		log.Fatalf("Error al obtener empleados: %v", err)
	}
	defer rows.Close()

	var employees []entities.Employee
	for rows.Next() {
		var employee entities.Employee
		if err := rows.Scan(&employee.ID, &employee.Name); err != nil {
			log.Printf("Error al escanear el empleado: %v", err)
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error al iterar sobre las filas: %w", err)
	}

	return employees
}

func (m *MySQL) Update(employee entities.Employee) error {
	query := "UPDATE employees SET name = ? WHERE id = ?"
	_, err := m.conn.ExecutePreparedQuery(query, employee.Name, employee.ID)
	if err != nil {
		log.Printf("Error al actualizar el empleado: %v", err)
		return err
	}
	log.Printf("Empleado actualizado correctamente: %d", employee.ID)
	return nil
}

func (m *MySQL) Delete(employee entities.Employee) error {
	query := "DELETE FROM employees WHERE id = ?"
	_, err := m.conn.ExecutePreparedQuery(query, employee.ID)
	if err != nil {
		log.Printf("Error al eliminar el empleado: %v", err)
		return err
	}
	return nil
}

func (m *MySQL) FindByID(id int) (*entities.Employee, error) {
	query := "SELECT * FROM employees WHERE id = ?"
	log.Printf("Ejecutando consulta: %s con ID: %d", query, id)
	rows, err := m.conn.FetchRows(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employee entities.Employee
	if rows.Next() {
		err = rows.Scan(&employee.ID, &employee.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("Empleado no encontrado")
			}
			return nil, err
		}
		return &employee, nil
	}

	return nil, fmt.Errorf("Empleado no encontrado")
}
