package router

import (
	"go-test-class-api/server"

	"github.com/gorilla/mux"
)

// BindRoutes bind all the app routes
func BindRoutes(s server.Server, r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	// Every version must have it's own BindRoutesVX
	BindRoutesV1(s, api)
}
