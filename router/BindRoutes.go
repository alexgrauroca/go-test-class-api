package router

import (
	"go-test-class-api/server"

	"github.com/gorilla/mux"
)

func BindRoutes(s server.Server, r *mux.Router) {
	BindRoutesV1(s, r)
}
