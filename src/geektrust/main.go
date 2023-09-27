package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var Courses = make(map[string]Course)
var Allotments = make(map[string]Allotment)

type Employee struct {
	name  string
	email string
}

type Course struct {
	offeringId string
	name       string
	instructor string
	date       time.Time
	minEmp     int
	maxEmp     int
	current    int
	status     Status
}

type Allotment struct {
	email               string
	courseOfferingId    string
	courseName          string
	instructor          string
	date                time.Time
	courseRegisrationId string
	status              Status
}

type Status int

const (
	ACCEPTED          Status = 1
	CANCEL_ACCEPTED   Status = 2
	CANCEL_REJECTED   Status = 3
	COURSE_FULL_ERROR Status = 4
	COURSE_CANCELED   Status = 5
)

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	// count := 0
	for scanner.Scan() {

		args := scanner.Text()
		argList := strings.Fields(args)
		// fmt.Println(argList)
		if argList[0] == "ADD-COURSE-OFFERING" {
			fmt.Println("Add course")
			minEmp, err := strconv.Atoi(argList[4])
			if err != nil {
				fmt.Println("Unable to parse min emp")
			}

			maxEmp, err := strconv.Atoi(argList[5])
			if err != nil {
				fmt.Println("Unable to parse max emp")
			}
			t, err := time.Parse("02012006", argList[3])
			if err != nil {
				fmt.Println("Unable to parse time")
			}
			c := Course{
				"",
				argList[1],
				argList[2],
				t,
				minEmp,
				maxEmp,
				0,
				0,
			}
			fmt.Println("Adding course", c.date.Format("02012006"))
			fmt.Println(addCourse(c))

		} else if argList[0] == "REGISTER" {
			fmt.Println("\nRegister course")
			name := strings.Split(argList[1], "@")
			emp := Employee{
				name[0],
				argList[1],
			}

			//get course from courseID
			if course, ok := Courses[argList[2]]; !ok {
				fmt.Println("Course not found")
				return
			} else {
				a := Allotment{
					emp.email,
					course.offeringId,
					course.name,
					course.instructor,
					course.date,
					"",
					0,
				}
				fmt.Print(registerCourse(a))
			}

			// fmt.Println(a)
			// fmt.Println("Registering")
			//

		} else if argList[0] == "ALLOT-COURSE" {
			fmt.Println("Allotment List")
			// allotCourse(argList[1])

		} else if argList[0] == "CANCEL" {
			fmt.Println("Cancelling course")
			// fmt.Print(cancelRegistration(argList[1]))
		}
	}

	fmt.Println("\n\n", Courses)

	// fmt.Println("\n\nallotments", Allotments)
}

func joinStrings(str ...string) string {
	return strings.Join(str, "-")
}

func addCourse(c Course) string {
	c.offeringId = joinStrings("OFFERING", strings.ToUpper(c.name), strings.ToUpper(c.instructor))
	Courses[c.offeringId] = c
	return c.offeringId
}

func registerCourse(a Allotment) (string, Status) {
	//REG-COURSE-<EMPLOYEE-NAME>-<COURSE-NAME>
	courseRegistrationId := joinStrings("REG-COURSE", strings.ToUpper(strings.Split(a.email, "@")[0]), strings.ToUpper(a.courseName))
	var status Status
	currentDate := time.Now()

	if c, ok := Courses[a.courseOfferingId]; !ok {
		fmt.Println("Course not found")
		return "", 0
	} else {
		if Courses[a.courseOfferingId].current > c.maxEmp {
			// fmt.Println("Course full")
			status = 4
		} else if Courses[a.courseOfferingId].current < c.minEmp && c.date.After(currentDate) {
			status = 1
			temp := Courses[a.courseOfferingId].current
			temp++
			Courses[a.courseOfferingId] = Course{
				c.offeringId,
				c.name,
				c.instructor,
				c.date,
				c.minEmp,
				c.maxEmp,
				temp,
				status,
			}
			fmt.Println("Updated Course status")

		} else if Courses[a.courseOfferingId].current < c.minEmp && c.date.Before(currentDate) {
			// fmt.Println("Course cancelled")
			status = 5
		}
	}
	a.status = status
	Allotments[a.courseOfferingId] = a

	return courseRegistrationId, status
}

func findCourse(allotmentMap map[string][]Allotment, courseRegistrationId string) Allotment {
	// courseName := strings.Split(courseRegistrationId, "-")[2]
	// for k, v := range Allotments {
	// 	cName := strings.Split(k, "-")[1]
	// 	if cName == courseName {
	// 		for _, j := range v {
	// 			fmt.Println(j)
	// 			if j.courseRegisrationId == courseRegistrationId {
	// 				return j
	// 			}
	// 		}
	// 	}
	// }
	return Allotment{}
}

func cancelRegistration(courseRegistrationId string) (string, Status) {
	// var allotment Allotment
	// if _, ok := Allotments[courseRegistrationId]; !ok {
	// 	fmt.Println("Course not found!")
	// 	return courseRegistrationId, 0
	// } else {
	// 	allotment = findCourse(Allotments, courseRegistrationId)
	// 	if allotment.date.Before(time.Now()) {
	// 		allotment.status = CANCEL_REJECTED
	// 	} else {
	// 		allotment.status = CANCEL_ACCEPTED
	// 	}
	// }

	return courseRegistrationId, 0
}

func allotCourse(courseOfferingId string) []Allotment {
	//sort map based on registration number
	// if courses, ok := Allotments[courseOfferingId]; !ok {
	// 	fmt.Println("Course not found")
	// } else {
	// 	fmt.Print("\n", courseOfferingId, ": ")
	// 	for i := range courses {
	// 		fmt.Println("Course Name:", courses[i].courseName)
	// 		fmt.Println("Instructor Name:", courses[i].instructor)
	// 		fmt.Println("Email Id:", courses[i].email)
	// 		fmt.Println("Date: ", courses[i].date.Format("02012006"))
	// 		fmt.Printf("Status: %v\n", courses[i].status)
	// 	}
	// 	fmt.Println("\n")
	// 	return courses
	// }
	return nil
}

//to do
// cancel / allotment logic
// how to process input from file -- done
// refactor to logic files
