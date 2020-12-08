package render

import (
	"encoding/json"
	"errors"
	"net/http"

	"order-manager/pkg/api/v3/errs"
)

const (
	statusError   = "error"
	statusSuccess = "success"
)

type BodyError struct {
	Status  string `json:"status" default:"error"`
	Message string `json:"message" example:"something went wrong"`
}

type BodySuccess struct {
	Status string      `json:"status" default:"success"`
	Data   interface{} `json:"data,omitempty"`
}

func Error(w http.ResponseWriter, err error) {
	respondError(w, resolveCode(err), err)
}

func OK(w http.ResponseWriter, data interface{}) {
	respondSuccess(w, http.StatusOK, data)
}

func Created(w http.ResponseWriter, data interface{}) {
	respondSuccess(w, http.StatusCreated, data)
}

func resolveCode(err error) int {
	if errors.Is(err, errs.ErrNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, errs.ErrNotValid) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func respondError(w http.ResponseWriter, code int, err error) {
	var r BodyError
	r.Status = statusError
	r.Message = err.Error()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(r)
}

func respondSuccess(w http.ResponseWriter, code int, data interface{}) {
	var r BodySuccess
	r.Status = statusSuccess
	if data != nil {
		r.Data = data
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(r)
}
