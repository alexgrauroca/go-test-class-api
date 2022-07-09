package models

import (
	"go-test-class-api/helpers"
	"time"
)

type Class struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Capacity  int       `json:"capacity"`
}

func (c *Class) NewId() error {
	id, err := helpers.GenerateId()

	if err != nil {
		return err
	}

	c.Id = id

	return nil
}
