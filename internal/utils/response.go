package utils

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Success int         `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type errorResponse struct {
	Success int    `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func Success(w http.ResponseWriter, status int, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(successResponse{
		Success: status,
		Message: message,
		Data:    data,
	})
}
func Error(w http.ResponseWriter, status int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := errorResponse{
		Success: status,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}

func Created(w http.ResponseWriter, data any) {
	type envelop struct {
		Data any `json:"data"`
	}
	Success(w, http.StatusCreated, &envelop{Data: data}, "Created successfully")
}

func BadRequest(w http.ResponseWriter, err error) {
	Error(w, http.StatusBadRequest, "Bad request", err)
}

func InternalError(w http.ResponseWriter, err error) {
	Error(w, http.StatusInternalServerError, "Internal server error", err)
}

func NotFound(w http.ResponseWriter) {
	Error(w, http.StatusNotFound, "Not found", nil)
}

func MethodNotAllowed(w http.ResponseWriter) {
	Error(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
}
func Unauthorized(w http.ResponseWriter, err error) {
	Error(w, http.StatusUnauthorized, "Unauthorized", err)
}
func Forbidden(w http.ResponseWriter) {
	Error(w, http.StatusForbidden, "Forbidden", nil)
}
