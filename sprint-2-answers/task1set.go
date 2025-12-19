package sprint_2_answers

func Intersect(setA, setB map[int]bool) map[int]bool {
	result := make(map[int]bool)

	if len(setA) > len(setB) {
		setA, setB = setB, setA
	}

	for k := range setA {
		if setB[k] {
			result[k] = true
		}
	}

	return result
}
