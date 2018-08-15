package jsonresponse

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeJSON struct {
	Field1 int64
	Field2 string
}

func testJSONHeaders(t *testing.T) {
	rr := httptest.NewRecorder()
	jsonHeaders(rr)
	if rr.Header().Get("Content-Type") != "application/json; charset=utf-8" {
		t.Fatalf(`Unexpected result, got: %s`, rr.Header().Get("Content-Type"))
	}
}
func TestOk(t *testing.T) {
	rr := httptest.NewRecorder()
	OK(rr, &fakeJSON{Field1: 1234, Field2: "test string"})

	if rr.Code != http.StatusOK {
		t.Fatalf(`Unexpected result, got: %d`, rr.Code)
	}

	if rr.Body.String() != `{"Field1":1234,"Field2":"test string"}` {
		t.Fatalf(`Unexpected result, got: %s`, rr.Body.String())
	}
}

func TestCreated(t *testing.T) {
	rr := httptest.NewRecorder()
	Created(rr, &fakeJSON{Field1: 1234, Field2: "test string"})

	if rr.Code != http.StatusCreated {
		t.Fatalf(`Unexpected result, got: %d`, rr.Code)
	}

	if rr.Body.String() != `{"Field1":1234,"Field2":"test string"}` {
		t.Fatalf(`Unexpected result, got: %s`, rr.Body.String())
	}
}

func TestBadRequest(t *testing.T) {
	rr := httptest.NewRecorder()
	BadRequest(rr, errors.New("this is a bad request"))

	if rr.Code != http.StatusBadRequest {
		t.Fatalf(`Unexpected result, got code: %d`, rr.Code)
	}

	if rr.Body.String() != `{"error_message":"this is a bad request"}` {
		t.Fatalf(`Unexpected result, got: %s`, rr.Body.String())
	}
}

func TestNotFound(t *testing.T) {
	rr := httptest.NewRecorder()
	NotFound(rr, errors.New("resource not found"))

	if rr.Code != http.StatusNotFound {
		t.Fatalf(`Unexpected result, got code: %d`, rr.Code)
	}

	if rr.Body.String() != `{"error_message":"resource not found"}` {
		t.Fatalf(`Unexpected result, got: %s`, rr.Body.String())
	}
}
