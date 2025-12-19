package school

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSchoolReport(t *testing.T) {
	students := []Student{
		{Name: "Sergey", Age: 24, Gender: "male", Height: 184},
		{Name: "Anton", Age: 23, Gender: "male", Height: 184},
		{Name: "Tolya", Age: 26, Gender: "male", Height: 163},
		{Name: "Anastasiya", Age: 16, Gender: "female", Height: 154},
		{Name: "Dmitriy", Age: 26, Gender: "male", Height: 181},
		{},
		{Name: "Ivan", Age: 17, Gender: "male", Height: 165},
		{Name: "Anna", Age: 18, Gender: "female", Height: 158},
		{Name: "Petr", Age: 19, Gender: "male", Height: 172},
		{Name: "Olga", Age: 16, Gender: "female", Height: 149},
		{Name: "Max", Age: 20, Gender: "male", Height: 181},
	}

	report := getSchoolReport(students)

	require.Equal(t, 7, report.BoysCount)
	require.Equal(t, 3, report.GirlsCount)
	require.Equal(t, 7, report.AdultsCount)

	require.Equal(t, map[string]int{
		"<150":    1,
		"150-160": 2,
		"160-170": 2,
		"170-180": 1,
		">=180":   4,
	}, report.HeightGroups)
}
