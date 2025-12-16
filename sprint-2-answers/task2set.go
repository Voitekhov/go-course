package sprint_2_answers

func FindDifference(setA, setB map[int]bool) map[int]bool {
	result := make(map[int]bool)

	for k := range setA {
		if !setB[k] {
			result[k] = true
		}
	}

	return result
}
