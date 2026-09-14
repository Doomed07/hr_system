package http

import (
	maincommands "HR_system/main_commands"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Handlers struct {
	Table *maincommands.TableEmployees
}

func NewHandlers(table *maincommands.TableEmployees) *Handlers {
	return &Handlers{
		Table: table,
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

	if err := h.Table.AddEmp(maincommands.Employee(employee)); err != nil {
		if errors.Is(err, maincommands.ErrIdAlreadyUsed) {
			h.writeErr(w, err, http.StatusConflict)
			return
		}
		h.writeErr(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	h.Output(w, employee)
}

func (h *Handlers) HandleAllEmployee(w http.ResponseWriter, r *http.Request) {
	tableDTO, err := h.Table.AllEmp()
	if err != nil {
		if errors.Is(err, maincommands.ErrTableIsEmpty) {
			h.writeErr(w, err, http.StatusBadRequest)
			return
		}
		h.writeErr(w, err, http.StatusInternalServerError)
		return
	}

	h.Output(w, tableDTO)
}

func (h *Handlers) HandelDelEmployee(w http.ResponseWriter, r *http.Request) {
	var idDTO IdDTO
	if err := json.NewDecoder(r.Body).Decode(&idDTO); err != nil {
		h.writeErr(w, err, http.StatusBadRequest)
		return
	}

	if err := h.Table.DelEmp(idDTO.ID); err != nil {
		if errors.Is(err, maincommands.ErrNotFoundEmployee) {
			h.writeErr(w, err, http.StatusBadRequest)
			return
		}
		h.writeErr(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
