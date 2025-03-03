package application

import "demo/src/employees/domain/entities"
import "demo/src/employees/domain"


type CreateEmployee struct {
	db domain.IEmployeeRepository
}

func NewCreateEmployee(db domain.IEmployeeRepository) *CreateEmployee {
	return &CreateEmployee{db: db}
}

func (ce *CreateEmployee) Execute(employee entities.Employee) {
	ce.db.Create(employee)
}


