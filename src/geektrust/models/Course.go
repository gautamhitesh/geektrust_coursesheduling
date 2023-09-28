package models

import "time"

type Course struct {
	OfferingId string
	Name       string
	Instructor string
	Date       time.Time
	MinEmp     int
	MaxEmp     int
	Current    int
	Status     Status
}
