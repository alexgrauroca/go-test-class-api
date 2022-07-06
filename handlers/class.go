package handlers

import (
	"encoding/json"
	"go-test-class-api/helpers"
	"go-test-class-api/models"
	"go-test-class-api/repository"
	"go-test-class-api/server"
	"net/http"
	"time"
)

type InsertClassRequest struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Capacity  int       `json:"capacity"`
}

func InsertClassHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var classRequest = InsertClassRequest{}

		if err := json.NewDecoder(r.Body).Decode(&classRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		c := models.Class{
			Name:      classRequest.Name,
			StartDate: classRequest.StartDate,
			EndDate:   classRequest.EndDate,
			Capacity:  classRequest.Capacity,
		}

		if err := c.NewId(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := repository.InsertClass(r.Context(), &c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helpers.HttpJsonResponse(w, c, http.StatusCreated)
	}
}
