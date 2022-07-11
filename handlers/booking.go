package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alexgrauroca/go-test-class-api/helpers"
	"github.com/alexgrauroca/go-test-class-api/models"
	"github.com/alexgrauroca/go-test-class-api/repository"
	"github.com/alexgrauroca/go-test-class-api/server"
)

type InsertBookingRequest struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

func InsertBookingHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		InsertBooking(w, r)
	}
}

func InsertBooking(w http.ResponseWriter, r *http.Request) {
	var bookingRequest = InsertBookingRequest{}

	if err := json.NewDecoder(r.Body).Decode(&bookingRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	booking := models.Booking{
		Name: bookingRequest.Name,
		Date: bookingRequest.Date,
	}

	if err := booking.NewId(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := repository.InsertBooking(r.Context(), &booking); err != nil {
		errorStatus := http.StatusInternalServerError

		if err.Error() == "No classes are available for the specified date" {
			errorStatus = http.StatusBadRequest
		}

		http.Error(w, err.Error(), errorStatus)
		return
	}

	helpers.HttpJsonResponse(w, booking, http.StatusCreated)
}
