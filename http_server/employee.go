package http_server

import "errors"

type Employee struct {
	ID       int
	FullName string
	Position string
}

func NewEmpoyee(id int, fullname, position string) Employee {
	return Employee{
		ID:       id,
		FullName: fullname,
		Position: position,
	}
}

var ErrNotFoundEmployee = errors.New("Not found employee with specified ID")
