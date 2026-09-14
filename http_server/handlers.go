package http_server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type EmployeesStorage interface {
	AddEmp(ctx context.Context, emp Employee) (int, error)
	AllEmp(ctx context.Context) ([]Employee, error)
	DelEmp(ctx context.Context, id int) error
}

type Handlers struct {
	Storage EmployeesStorage
}

func NewHandlers(storege EmployeesStorage) *Handlers {
	return &Handlers{
		Storage: storege,
	}
}

func (h *Handlers) writeErr(w http.ResponseWriter, err error, code int) {
	errDTO := ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
	http.Error(w, errDTO.ToStr(), code)
}

func (h *Handlers) Output(w http.ResponseWriter, t any) {
	b, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		h.writeErr(w, err, http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		h.writeErr(w, err, http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) HandleNewEmployee(w http.ResponseWriter, r *http.Request) {
	var employee EmpDTO
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		h.writeErr(w, err, http.StatusBadRequest)
		return
	}

	id, err := h.Storage.AddEmp(r.Context(), Employee(employee))
	if err != nil {
		h.writeErr(w, err, http.StatusInternalServerError)
		return
	}

	employee.ID = id

	w.WriteHeader(http.StatusCreated)
	h.Output(w, employee)
}

func (h *Handlers) HandleAllEmployee(w http.ResponseWriter, r *http.Request) {
	employees, err := h.Storage.AllEmp(r.Context())
	if err != nil {
		h.writeErr(w, err, http.StatusInternalServerError)
		return
	}

	resp := make([]EmpDTO, 0, len(employees))
	for _, employee := range employees {
		resp = append(resp, EmpDTO(employee))
	}

	h.Output(w, resp)
}

func (h *Handlers) HandelDelEmployee(w http.ResponseWriter, r *http.Request) {
	var idDTO IdDTO
	if err := json.NewDecoder(r.Body).Decode(&idDTO); err != nil {
		h.writeErr(w, err, http.StatusBadRequest)
		return
	}

	if err := h.Storage.DelEmp(r.Context(), idDTO.ID); err != nil {
		if errors.Is(err, ErrNotFoundEmployee) {
			h.writeErr(w, err, http.StatusNotFound)
			return
		}
		h.writeErr(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
