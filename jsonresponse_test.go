package jsonresponse

import (
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
