package sprint_2_answers

func SymmetricDifference(setA, setB map[int]bool) map[int]bool {
	result := make(map[int]bool)

	for k := range setA {
		if !setB[k] {
			result[k] = true
		}
	}

	for k := range setB {
		if !setA[k] {
			result[k] = true
		}
	}

	return result
}
