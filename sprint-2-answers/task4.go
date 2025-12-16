package sprint_2_answers

func SwapMap(numbers map[string]int) map[int]string {
	result := make(map[int]string)
	for key, value := range numbers {
		result[value] = key
	}
	return result
}
