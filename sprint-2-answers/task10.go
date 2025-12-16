package sprint_2_answers

func BuildLetterMap(message string) map[rune]int {
	letter := make(map[rune]int)

	for _, r := range message {
		if r == ' ' {
			continue
		}
		letter[r]++
	}

	return letter
}
