package maincommands

import (
	"sync"
)

type TableEmployees struct {
	Table map[int]Employee
	mtx   sync.RWMutex
}

func NewTableEmployees() *TableEmployees {
	return &TableEmployees{
		Table: make(map[int]Employee),
	}
}

func (t *TableEmployees) AddEmp(employee Employee) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	if _, ok := t.Table[employee.ID]; ok {
		return ErrIdAlreadyUsed
	}

	t.Table[employee.ID] = employee

	return nil
}

func (t *TableEmployees) AllEmp() (map[int]Employee, error) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	if t.Table == nil {
		return nil, ErrTableIsEmpty
	}
	return t.Table, nil
}

func (t *TableEmployees) DelEmp(id int) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	_, ok := t.Table[id]
	if !ok {
		return ErrNotFoundEmployee
	}

	delete(t.Table, id)
	return nil
}
