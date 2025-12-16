package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountUnique(t *testing.T) {

	testTable := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{},
			expected: 0,
		},
		{
			nums:     []int{1, 2, 2, 3, 4, 4, 4, 5},
			expected: 5,
		},
		{
			nums:     []int{1, 1, 1, 1},
			expected: 1,
		},
		{
			nums:     []int{1, 2, 3, 4},
			expected: 4,
		},
	}

	for _, tc := range testTable {
		got := CountUnique(tc.nums)
		require.Equal(t, tc.expected, got)
	}
}
