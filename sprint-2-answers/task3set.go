package sprint_2_answers

func SetUnion(setA, setB map[string]bool) map[string]bool {
	result := make(map[string]bool)

	for k := range setA {
		result[k] = true
	}

	for k := range setB {
		result[k] = true
	}

	return result
}
