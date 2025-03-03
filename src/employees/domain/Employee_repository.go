package domain

import "demo/src/employees/domain/entities"

type IEmployeeRepository interface {
	GetAll() []entities.Employee
	Create(employee entities.Employee)
	Update(employee entities.Employee) error
	Delete(employee entities.Employee) error
	FindByID(id int) (*entities.Employee, error)
}
