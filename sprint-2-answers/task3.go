package sprint_2_answers

func FilteringByValue(items map[string]int, min int) map[string]int {
	result := make(map[string]int)

	for key, value := range items {
		if value > min {
			result[key] = value
		}
	}

	return result
}
