package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeAndSum(t *testing.T) {

	testTable := []struct {
		mapA     map[string]int
		mapB     map[string]int
		expected map[string]int
	}{
		{
			mapA: map[string]int{
				"apple":  5,
				"banana": 2},
			mapB: map[string]int{
				"banana": 3,
				"orange": 4},
			expected: map[string]int{
				"apple":  5,
				"banana": 5,
				"orange": 4},
		},
		{
			mapA: map[string]int{
				"apple":  5,
				"banana": 2,
				"orange": 4},
			mapB: map[string]int{
				"apple":  5,
				"banana": 2,
				"orange": 4},
			expected: map[string]int{
				"apple":  10,
				"banana": 4,
				"orange": 8},
		},
		{
			mapA:     map[string]int{},
			mapB:     map[string]int{},
			expected: map[string]int{},
		},
	}

	for _, tc := range testTable {
		got := MergeAndSum(tc.mapA, tc.mapB)
		require.Equal(t, tc.expected, got)
	}
}
