package handlers

import (
	"go-test-class-api/server"
	"net/http"
)

func InsertBookingHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		InsertBooking(w, r)
	}
}

func InsertBooking(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusInternalServerError)
}
