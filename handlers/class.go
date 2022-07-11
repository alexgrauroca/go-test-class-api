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

type InsertClassRequest struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Capacity  int       `json:"capacity"`
}

type ClassesResponse struct {
	Classes []*models.Class `json:"classes"`
}

func InsertClassHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		InsertClass(w, r)
	}
}

func InsertClass(w http.ResponseWriter, r *http.Request) {
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
		errorStatus := http.StatusInternalServerError

		if err.Error() == "There can only be one class a day" {
			errorStatus = http.StatusBadRequest
		}

		http.Error(w, err.Error(), errorStatus)
		return
	}

	helpers.HttpJsonResponse(w, c, http.StatusCreated)
}

func GetClassesHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GetClasses(w, r)
	}
}

func GetClasses(w http.ResponseWriter, r *http.Request) {
	var filters = models.ClassFilter{}
	classes, err := repository.GetClasses(r.Context(), &filters)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(classes) == 0 {
		http.Error(w, "Classes not found", http.StatusNotFound)
		return
	}

	helpers.HttpJsonResponse(w, ClassesResponse{Classes: classes}, http.StatusOK)
}
