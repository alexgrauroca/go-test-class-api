package models

import (
	"time"
)

type ClassFilter struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Date      time.Time `json:"date"`
}
