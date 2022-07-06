package router

import (
	"go-test-class-api/server"
	"net/http"

	"github.com/gorilla/mux"
)

func addRoutesV1(s server.Server, r *mux.Router) {
	r.HandleFunc("/classes", handlers.InsertPostHandler(s)).Methods(http.MethodPost)
}
