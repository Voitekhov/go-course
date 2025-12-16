package sprint_2_answers

func CheckSubSet(setA, setB map[int]bool) bool {

	if len(setA) > len(setB) {
		return false
	}
	for k := range setA {
		if !setB[k] {
			return false
		}
	}
	return true
}
