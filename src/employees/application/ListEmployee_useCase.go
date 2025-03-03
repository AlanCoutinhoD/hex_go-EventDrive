package application

import "demo/src/employees/domain/entities"
import "demo/src/employees/domain"

type ListEmployee struct {
	db domain.IEmployeeRepository
}

func NewListEmployee(db domain.IEmployeeRepository) *ListEmployee {
	return &ListEmployee{db: db}
}

func (le *ListEmployee) Execute() []entities.Employee {
	return le.db.GetAll()
}


