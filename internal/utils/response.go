package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type successResponse struct {
	Success int         `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type errorResponse struct {
	Success int            `json:"success"`
	Message string         `json:"message"`
	Error   string         `json:"error,omitempty"`
	Errors  []ValidatorErr `json:"errors,omitempty"`
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
func Error(w http.ResponseWriter, status int, message string, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := errorResponse{
		Success: status,
		Message: message,
	}
	// Handle different types of errors
	switch v := err.(type) {
	case error:
		response.Error = v.Error()
	case []ValidatorErr:
		response.Errors = v
	case string:
		response.Error = v
	case nil:
		// No error data
	default:
		// Handle other types if needed
		response.Error = fmt.Sprintf("%v", v)
	}

	json.NewEncoder(w).Encode(response)
}

func Created(w http.ResponseWriter, data any) {
	type envelop struct {
		Data any `json:"data"`
	}
	Success(w, http.StatusCreated, &envelop{Data: data}, "Created successfully")
}

func BadRequest(w http.ResponseWriter, message string, err error) {
	Error(w, http.StatusBadRequest, message, err)
}
func ValidationError(w http.ResponseWriter, message string, validationErrors []ValidatorErr) {
	Error(w, http.StatusBadRequest, message, validationErrors)
}

func InternalError(w http.ResponseWriter, err error, message ...string) {
	msg := "Internal server error"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Error(w, http.StatusInternalServerError, msg, err)
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

func ReadJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteJsonError(w http.ResponseWriter, status int, message string) error {
	type envelop struct {
		Error string `json:"error"`
	}
	return writeJson(w, status, &envelop{Error: message})
}

func writeJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
