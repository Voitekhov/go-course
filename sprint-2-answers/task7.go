package sprint_2_answers

func InvertBoolMap(mapToInvert map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for key, value := range mapToInvert {
		result[key] = !value
	}
	return result
}
