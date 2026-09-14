package maincommands

import "errors"

var ErrIdAlreadyUsed = errors.New("ID already has been used")
var ErrTableIsEmpty = errors.New("Table is empty")
var ErrNotFoundEmployee = errors.New("Not found employee with specified ID")
