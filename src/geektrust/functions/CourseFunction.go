package functions

import (
	"fmt"
	. "geektrust/models"
	"strings"
)

var Courses = make(map[string]Course)

func AddCourse(c Course) string {
	// OFFERING-COURSENAME-INSTRUCTOR
	c.OfferingId = joinStrings("OFFERING", strings.ToUpper(c.Name), strings.ToUpper(c.Instructor))
	Courses[c.OfferingId] = c
	return c.OfferingId
}

func joinStrings(str ...string) string {
	return strings.Join(str, "-")
}

func CancelRegistration(courseRegistrationId string) (string, Status) {
	// flag := false
	for k, v := range Allotments {
		if strings.Split(k, "-")[1] == strings.Split(courseRegistrationId, "-")[3] {
			// fmt.Println("Course found")
			temp := v
			for key, value := range v {
				if value.CourseRegisrationId == courseRegistrationId {
					t := temp[key]
					if value.Status == ACCEPTED {
						t.Status = CANCEL_REJECTED
						t.Instructor = "Updated cancelled"
						Allotments[k][key] = t
						return courseRegistrationId, temp[key].Status
					} else {
						t.Status = CANCEL_ACCEPTED
						t.Instructor = "Updated updated"
						Allotments[k][key] = t
						return courseRegistrationId, temp[key].Status
					}
				}
			}
			// fmt.Println("Registration not found", flag, courseRegistrationId)
		}
	}
	return courseRegistrationId, -1
}

func RegisterCourse(a Allotment) (string, Status) {
	//REG-COURSE-<EMPLOYEE-NAME>-<COURSE-NAME>
	courseRegistrationId := joinStrings("REG-COURSE", strings.ToUpper(strings.Split(a.Email, "@")[0]), strings.ToUpper(a.CourseName))
	var status Status
	var astatus Status
	var c Course
	var temp int
	if _, ok := Courses[a.CourseOfferingId]; !ok {
		fmt.Println("Course not found")
		return "", -1
	} else {
		c = Courses[a.CourseOfferingId]
		if c.Current >= c.MaxEmp {
			status = COURSE_FULL_ERROR
		} else if c.Current < c.MaxEmp {
			status = ACCEPTED
			astatus = CONFIRMED
			temp = c.Current + 1
		}
		Courses[a.CourseOfferingId] = Course{
			OfferingId: c.OfferingId,
			Name:       c.Name,
			Instructor: c.Instructor,
			Date:       c.Date,
			MinEmp:     c.MinEmp,
			MaxEmp:     c.MaxEmp,
			Current:    temp,
			Status:     status,
		}
	}
	a.CourseRegisrationId = courseRegistrationId
	a.Status = astatus
	Allotments[a.CourseOfferingId] = append(Allotments[a.CourseOfferingId], a)
	if status == ACCEPTED {
		return courseRegistrationId, status
	} else {
		return "", status
	}

}
