package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupByLength(t *testing.T) {

	testTable := []struct {
		words    []string
		expected map[int][]string
	}{
		{
			words: []string{"go", "java", "rust", "c", "js"},
			expected: map[int][]string{
				1: {"c"},
				2: {"go", "js"},
				4: {"java", "rust"},
			},
		},
		{
			words:    []string{},
			expected: map[int][]string{},
		},
		{
			words: []string{"a", "b", "c"},
			expected: map[int][]string{
				1: {"a", "b", "c"},
			},
		},
	}

	for _, tc := range testTable {
		got := GroupByLength(tc.words)
		require.Equal(t, tc.expected, got)
	}
}
