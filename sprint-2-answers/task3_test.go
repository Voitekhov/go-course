package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilteringByValue(t *testing.T) {

	testTable := []struct {
		min      int
		items    map[string]int
		expected map[string]int
	}{
		{
			items: map[string]int{
				"apple":  5,
				"banana": 1,
				"orange": 10,
			},
			min: 3,
			expected: map[string]int{
				"apple":  5,
				"orange": 10,
			},
		},
		{
			items: map[string]int{
				"apple":  5,
				"banana": 1,
				"orange": 10,
			},
			min: 0,
			expected: map[string]int{
				"apple":  5,
				"banana": 1,
				"orange": 10,
			},
		},
		{
			items:    map[string]int{},
			min:      10,
			expected: map[string]int{},
		},
	}

	for _, tc := range testTable {
		got := FilteringByValue(tc.items, tc.min)
		require.Equal(t, tc.expected, got)
	}
}
