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

// BadRequest return 400 Bad Request response
func BadRequest(w http.ResponseWriter, err error) {
	jsonHeaders(w)
	w.WriteHeader(http.StatusBadRequest)
}

// NotFound return 404 Not Found response
func NotFound(w http.ResponseWriter, err error) {
	jsonHeaders(w)
	w.WriteHeader(http.StatusNotFound)
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
