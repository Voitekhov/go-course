package sprint_2_answers

func CountWords(words []string) map[string]int {
	result := make(map[string]int)

	for _, w := range words {
		result[w]++
	}

	return result
}
