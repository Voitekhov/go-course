package sprint_2_answers

func Intersect(bigSet, smallSet map[int]bool) map[int]bool {
	result := make(map[int]bool)

	if len(bigSet) > len(smallSet) {
		bigSet, smallSet = smallSet, bigSet
	}

	for k := range bigSet {
		if smallSet[k] {
			result[k] = true
		}
	}

	return result
}
