package school

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSchoolReport(t *testing.T) {
	students := []Student{
		{Name: "Sergey", Age: 24, Gender: Male, Height: 184},
		{Name: "Anton", Age: 23, Gender: Male, Height: 184},
		{Name: "Tolya", Age: 26, Gender: Male, Height: 163},
		{Name: "Anastasiya", Age: 16, Gender: Female, Height: 154},
		{Name: "Dmitriy", Age: 26, Gender: Male, Height: 181},
		{},
		{Name: "Ivan", Age: 17, Gender: Male, Height: 165},
		{Name: "Anna", Age: 18, Gender: Female, Height: 158},
		{Name: "Petr", Age: 19, Gender: Male, Height: 172},
		{Name: "Olga", Age: 16, Gender: Female, Height: 149},
		{Name: "Max", Age: 20, Gender: Male, Height: 181},
	}

	report := getSchoolReport(students)

	require.Equal(t, 7, report.BoysCount)
	require.Equal(t, 3, report.GirlsCount)
	require.Equal(t, 7, report.AdultsCount)

	require.Equal(t, map[HeightGroup]int{
		HeightLess150:    1,
		Height150to160:   2,
		Height160to170:   2,
		Height170to180:   1,
		Height180AndMore: 4,
	}, report.HeightGroups)
}
