package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountWords(t *testing.T) {

	testTable := []struct {
		words    []string
		expected map[string]int
	}{
		{
			words:    []string{},
			expected: map[string]int{},
		},
		{
			words: []string{"go", "java", "go", "python", "go", "java"},
			expected: map[string]int{
				"go":     3,
				"java":   2,
				"python": 1,
			},
		},
		{
			words: []string{"a", "b", "c"},
			expected: map[string]int{
				"a": 1,
				"b": 1,
				"c": 1,
			},
		},
	}

	for _, tc := range testTable {
		got := CountWords(tc.words)
		require.Equal(t, tc.expected, got)
	}
}
