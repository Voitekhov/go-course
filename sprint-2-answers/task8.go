package sprint_2_answers

func RemoveElementsByConditions(ratings map[int]int, limit int) map[int]int {
	result := make(map[int]int)
	for key, value := range ratings {
		if value < limit {
			delete(ratings, key)
		}
	}
	return result
}
