package models

import "time"

type Allotment struct {
	Email               string
	CourseOfferingId    string
	CourseName          string
	Instructor          string
	Date                time.Time
	CourseRegisrationId string
	Status              Status
}
