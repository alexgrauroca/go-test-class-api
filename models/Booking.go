package models

import (
	"go-test-class-api/helpers"
	"time"
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
