package testhelpers

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

// TestRequest test an http request
func TestRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.NewRouter().ServeHTTP(rr, req)

	return rr
}
