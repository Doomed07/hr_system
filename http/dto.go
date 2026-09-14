package http

import (
	"encoding/json"
	"time"
)

type EmpDTO struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Position string `json:"position"`
}

type IdDTO struct {
	ID int `json:"id"`
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToStr() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}
	return string(b)
}
