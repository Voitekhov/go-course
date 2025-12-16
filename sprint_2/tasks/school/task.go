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
		HeightGroups: map[string]int{
			"<150":    0,
			"150-160": 0,
			"160-170": 0,
			"170-180": 0,
			">=180":   0,
		},
	}
	for _, student := range students {

		if student.Gender == "male" {
			report.BoysCount++
		} else if student.Gender == "female" {
			report.GirlsCount++
		}

		if student.Age >= 18 {
			report.AdultsCount++
		}
		switch {
		case student.Height <= 0:
			continue
		case student.Height < 150:
			report.HeightGroups["<150"]++
		case student.Height >= 150 && student.Height < 160:
			report.HeightGroups["150-160"]++
		case student.Height >= 160 && student.Height < 170:
			report.HeightGroups["160-170"]++
		case student.Height >= 170 && student.Height < 180:
			report.HeightGroups["170-180"]++
		default:
			report.HeightGroups[">=180"]++
		}
	}
	return report
}

type Report struct {
	HeightGroups map[string]int
	BoysCount    int
	GirlsCount   int
	AdultsCount  int
}

type Student struct {
	Name   string
	Age    int
	Gender string
	Height int
}
