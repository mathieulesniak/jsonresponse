package jsonresponse

import (
	"encoding/json"
	"log"
	"net/http"
)

// OK return 200 OK reponse
func OK(w http.ResponseWriter, result interface{}) {
	jsonHeaders(w)
	w.WriteHeader(http.StatusOK)

	w.Write(jsonBody(result))
}

// Created return 201 Created response
func Created(w http.ResponseWriter, result interface{}) {
	jsonHeaders(w)
	w.WriteHeader(http.StatusCreated)

	w.Write(jsonBody(result))
}

// NoContent return 204 No Content response
func NoContent(w http.ResponseWriter) {
	jsonHeaders(w)
	w.WriteHeader(http.StatusNoContent)
}

// BadRequest return 400 Bad Request response
func BadRequest(w http.ResponseWriter, err error) {
	jsonHeaders(w)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonError(err))
}

// Unauthorized return a 401 Unauthorized response
func Unauthorized(w http.ResponseWriter, err error) {
	jsonHeaders(w)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(jsonError(err))
}

// Forbidden return a 403 Forbidden response
func Forbidden(w http.ResponseWriter, err error) {
	jsonHeaders(w)
	w.WriteHeader(http.StatusForbidden)
	w.Write(jsonError(err))
}

// NotFound return 404 Not Found response
func NotFound(w http.ResponseWriter, err error) {
	jsonHeaders(w)
	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonError(err))
}

// InternalServerError return 500 Internal Server Error response
func InternalServerError(w http.ResponseWriter, err error) {
	jsonHeaders(w)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonError(err))
}

func jsonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func jsonError(err error) []byte {
	type errorMsg struct {
		ErrorMessage string `json:"error_message"`
	}
	errorStruct := errorMsg{ErrorMessage: err.Error()}
	encodedError, _ := json.Marshal(errorStruct)

	return encodedError
}

func jsonBody(r interface{}) []byte {
	body, err := json.Marshal(r)
	if err != nil {
		log.Println("Failed encoding JSON", err)
		return []byte("")
	}

	return body
}
