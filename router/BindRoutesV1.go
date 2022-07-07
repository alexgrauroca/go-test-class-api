package router

import (
	"go-test-class-api/server"

	"github.com/gorilla/mux"
)

// BindRoutesV1 Bind the app routes of version 1
func BindRoutesV1(s server.Server, r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()

	// The middlewers will be setted inside the next func, like the endpoints are setted inside addRoutesV1.
	// addMiddleweresV1(s, v1)
	addRoutesV1(s, v1)
}
