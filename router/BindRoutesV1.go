package router

import (
	"go-test-class-api/server"

	"github.com/gorilla/mux"
)

func BindRoutesV1(s server.Server, r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()

	addRoutesV1(s, api)
}
