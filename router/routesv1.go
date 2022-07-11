package router

import (
	"go-test-class-api/handlers"
	"go-test-class-api/server"
	"net/http"

	"github.com/gorilla/mux"
)

// addRoutesV1 Add v1 endpoints
func addRoutesV1(s server.Server, r *mux.Router) {
	r.HandleFunc("/bookings", handlers.InsertBookingHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/classes", handlers.GetClassesHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/classes", handlers.InsertClassHandler(s)).Methods(http.MethodPost)
}
