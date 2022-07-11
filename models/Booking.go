package models

import (
	"time"

	"github.com/alexgrauroca/go-test-class-api/helpers"
)

type Booking struct {
	Id   string    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

func (booking *Booking) NewId() error {
	id, err := helpers.GenerateId()

	if err != nil {
		return err
	}

	booking.Id = id

	return nil
}
