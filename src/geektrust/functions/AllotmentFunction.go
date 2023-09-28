package functions

import (
	"fmt"
	. "geektrust/models"
	"sort"
)

var Allotments = make(map[string][]Allotment)

func AllotCourse(courseOfferingId string) []Allotment {
	CourseAllotment(courseOfferingId)

	if course, ok := Allotments[courseOfferingId]; !ok {
		return nil
	} else {
		sort.Slice(course[:], func(i, j int) bool {
			return course[i].CourseRegisrationId < course[j].CourseRegisrationId
		})
		for _, i := range course {
			if i.Status == ACCEPTED {
				if Courses[courseOfferingId].Current < Courses[courseOfferingId].MinEmp {
					i.Status = COURSE_CANCELED
				} else {
					i.Status = CONFIRMED
				}

				fmt.Printf("%v %v %v %v %v %v %v\n", i.CourseRegisrationId, i.Email, i.CourseOfferingId, i.CourseName, i.Instructor, i.Date.Format("02012006"), i.Status.ToString())
			}
		}
		return course
	}
}
