package application

import (
	"demo/src/employees/domain"
	"demo/src/employees/domain/entities"
	"strconv"
)

type ListEmployeeById struct {
	repository domain.IEmployeeRepository
}

func NewListEmployeeById(repository domain.IEmployeeRepository) *ListEmployeeById {
	return &ListEmployeeById{repository: repository}
}

func (le *ListEmployeeById) Execute(id string) (*entities.Employee, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return le.repository.FindByID(intID)
} 