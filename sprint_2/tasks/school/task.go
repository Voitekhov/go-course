package school

/*
В школе учится множество учеников. У ученика есть имя, возраст, пол, вес и рост.
Нужно сделать сводку по школе со следующей информацией:
1. Разбивка учеников по росту:
	< 150
	[150-160)
	[160-170)
	[170-180)
	>= 180
2. Кол-во мальчиков и девочек
3. кол-во учеников, которым уже 18 лет и больше.

Задания:
 1. Спроектировать struct для ученика и отчета
 2. Решить задачу
 3. Написать тесты
*/

func getSchoolReport(students []Student) Report {
	report := Report{
		HeightGroups: map[HeightGroup]int{
			HeightLess150:    0,
			Height150to160:   0,
			Height160to170:   0,
			Height170to180:   0,
			Height180AndMore: 0,
		},
	}
	for _, student := range students {

		if student.Gender == Male {
			report.BoysCount++
		} else if student.Gender == Female {
			report.GirlsCount++
		}

		if student.Age >= 18 {
			report.AdultsCount++
		}
		switch {
		case student.Height <= 0:
			continue
		case student.Height < 150:
			report.HeightGroups[HeightLess150]++
		case student.Height >= 150 && student.Height < 160:
			report.HeightGroups[Height150to160]++
		case student.Height >= 160 && student.Height < 170:
			report.HeightGroups[Height160to170]++
		case student.Height >= 170 && student.Height < 180:
			report.HeightGroups[Height170to180]++
		default:
			report.HeightGroups[Height180AndMore]++
		}
	}
	return report
}

type Report struct {
	HeightGroups map[HeightGroup]int
	BoysCount    int
	GirlsCount   int
	AdultsCount  int
}

type Student struct {
	Name   string
	Age    int
	Gender Gender
	Height int
}

type Gender string

var Male = Gender("male")
var Female = Gender("female")

type HeightGroup string

const (
	HeightLess150    HeightGroup = "<150"
	Height150to160   HeightGroup = "150-160"
	Height160to170   HeightGroup = "160-170"
	Height170to180   HeightGroup = "170-180"
	Height180AndMore HeightGroup = ">=180"
)
