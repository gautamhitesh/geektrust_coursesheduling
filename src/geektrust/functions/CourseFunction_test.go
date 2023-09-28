package functions

import (
	. "geektrust/models"
	"reflect"
	"testing"
	"time"
)

func TestAddcourses(t *testing.T) {
	date, _ := time.Parse("02012006", "05062023")
	c := Course{
		OfferingId: "",
		Name:       "DATASCIENCE",
		Instructor: "BOB",
		Date:       date,
		MinEmp:     5,
		MaxEmp:     10,
		Current:    0,
	}
	result := AddCourse(c)
	want := "OFFERING-DATASCIENCE-BOB"
	if result != want {
		t.Errorf("Unable to add course %q %q", result, want)
	}
}

func TestAddCourse(t *testing.T) {
	tests := []struct {
		name string
		args Course
		want string
	}{

		// TODO: Add test cases.
		{"Test Course Add",
			Course{
				Instructor: "BOB",
				Name:       "DATASCIENCE",
			},
			"OFFERING-DATASCIENCE-BOB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddCourse(tt.args); got != tt.want {
				t.Errorf("AddCourse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_joinStrings(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test String join", args{str: []string{"Hello", "World"}}, "Hello-World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joinStrings(tt.args.str...); got != tt.want {
				t.Errorf("joinStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCourses(t *testing.T) {
	tests := []struct {
		name string
		want map[string]Course
	}{
		// TODO: Add test cases.
		{"Test Get All Courses",
			map[string]Course{"OFFERING-DATASCIENCE-BOB": {
				Instructor: "BOB",
				Name:       "DATASCIENCE",
				OfferingId: "OFFERING-DATASCIENCE-BOB",
				Date:       time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				MinEmp:     0,
				MaxEmp:     0,
				Status:     0,
				Current:    0,
			},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCourses(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCourses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourseAllotment(t *testing.T) {
	type args struct {
		coureseOfferingId string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CourseAllotment(tt.args.coureseOfferingId)
		})
	}
}

func TestCancelRegistration(t *testing.T) {
	type args struct {
		courseRegistrationId string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CancelRegistration(tt.args.courseRegistrationId)
			if got != tt.want {
				t.Errorf("CancelRegistration() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CancelRegistration() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegisterCourse(t *testing.T) {
	type args struct {
		a Allotment
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RegisterCourse(tt.args.a)
			if got != tt.want {
				t.Errorf("RegisterCourse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegisterCourse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
