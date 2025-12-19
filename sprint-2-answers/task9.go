package sprint_2_answers

func CountUnique(nums []int) int {
	seen := make(map[int]bool)

	for _, n := range nums {
		seen[n] = true
	}

	return len(seen)
}
