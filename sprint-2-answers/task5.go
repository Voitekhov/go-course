package sprint_2_answers

func GroupByLength(words []string) map[int][]string {
	result := make(map[int][]string)
	for _, w := range words {
		l := len(w)
		result[l] = append(result[l], w)
	}
	return result
}
