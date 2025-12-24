package sprint_2_answers

func wordCountByWord(words []string) map[string]int {
	result := make(map[string]int)

	for _, w := range words {
		result[w]++
	}

	return result
}
