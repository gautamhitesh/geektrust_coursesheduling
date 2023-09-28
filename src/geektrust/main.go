package main

import (
	"bufio"
	"fmt"
	. "geektrust/functions"
	. "geektrust/models"
	"os"
	"strconv"
	"strings"
	"time"
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

	for scanner.Scan() {

		args := scanner.Text()
		argList := strings.Fields(args)
		// fmt.Println(argList)
		if argList[0] == "ADD-COURSE-OFFERING" {
			// fmt.Println("Add course")
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
				OfferingId: "",
				Name:       argList[1],
				Instructor: argList[2],
				Date:       t,
				MinEmp:     minEmp,
				MaxEmp:     maxEmp,
				Current:    0,
				Status:     1,
			}

			fmt.Println(AddCourse(c))

		} else if argList[0] == "REGISTER" {
			// fmt.Println("\nRegister course")
			name := strings.Split(argList[1], "@")
			emp := Employee{
				Name:  name[0],
				Email: argList[1],
			}
			//get course from courseID
			if course, ok := Courses[argList[2]]; !ok {
				fmt.Println("Course not found")
				return
			} else {
				a := Allotment{
					Email:            emp.Email,
					CourseOfferingId: course.OfferingId,
					CourseName:       course.Name,
					Instructor:       course.Instructor,
					Date:             course.Date,
				}

				registerationId, status := RegisterCourse(a)
				if len(registerationId) == 0 {
					fmt.Printf("%v\n", status.ToString())
				} else {
					fmt.Printf("%v %v\n", registerationId, status.ToString())
				}
			}

		} else if argList[0] == "ALLOT" {
			// fmt.Println("Allotment List")
			AllotCourse(argList[1])

		} else if argList[0] == "CANCEL" {
			// fmt.Println("Cancelling course")
			reg, status := CancelRegistration(argList[1])
			fmt.Printf("%v %v\n", reg, status.ToString())
		}
	}
}
