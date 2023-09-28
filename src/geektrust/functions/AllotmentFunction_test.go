package functions

import (
	. "geektrust/models"
	"reflect"
	"testing"
)

func TestAllotCourse(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []Allotment
	}{
		// TODO: Add test cases.
		{"Test1", "OFFERING-DATASCIENCE-BOB", []Allotment{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllotCourse(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllotCourse() = %v, want %v", got, tt.want)
			}
		})
	}
}
