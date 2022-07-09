package models

import (
	"errors"
	"time"
)

type Booking struct {
	Id   string    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

func (booking *Booking) NewId() error {
	return errors.New("Not implemented")
}
