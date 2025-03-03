package entities

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	
}

func NewEmployee(id int, name string) *Employee {
	return &Employee{ID: id, Name: name}
}


func (e *Employee) GetName() string {
	return e.Name
}

func (e *Employee) SetName(name string) {
	e.Name = name
}

func (e *Employee) GetID() int {
	return e.ID
}




