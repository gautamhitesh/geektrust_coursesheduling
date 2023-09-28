package models

type Status int

const (
	CONFIRMED Status = iota
	ACCEPTED
	CANCEL_ACCEPTED
	CANCEL_REJECTED
	COURSE_FULL_ERROR
	COURSE_CANCELED
)

func (status Status) ToString() string {
	return [...]string{"CONFIRMED", "ACCEPTED", "CANCEL_ACCEPTED",
		"CANCEL_REJECTED",
		"COURSE_FULL_ERROR",
		"COURSE_CANCELED"}[status]
}
