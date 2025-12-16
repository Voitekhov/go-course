package sprint_2_answers

func MergeAndSum(mapA, mapB map[string]int) map[string]int {
	result := make(map[string]int)

	for k, v := range mapA {
		result[k] = v
	}

	for k, v := range mapB {
		result[k] += v
	}

	return result
}
