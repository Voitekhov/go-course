package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindDifference(t *testing.T) {

	testTable := []struct {
		setA     map[int]bool
		setB     map[int]bool
		expected map[int]bool
	}{
		{
			setA: map[int]bool{
				10: true, 20: true, 30: true, 40: true,
			},
			setB: map[int]bool{
				20: true, 40: true,
			},
			expected: map[int]bool{
				10: true, 30: true,
			},
		},
		{
			setA:     map[int]bool{},
			setB:     map[int]bool{},
			expected: map[int]bool{},
		},
		{
			setA: map[int]bool{
				1: true, 2: true, 3: true, 5: true,
			},
			setB: map[int]bool{
				10: true, 4: true, 11: true, 6: true,
			},
			expected: map[int]bool{1: true, 2: true, 3: true, 5: true},
		},
		{
			setA: map[int]bool{},
			setB: map[int]bool{
				3: true, 4: true, 5: true, 6: true,
			},
			expected: map[int]bool{},
		},
		{
			setA: map[int]bool{
				1: true, 2: true, 3: true, 5: true,
			},
			setB:     map[int]bool{},
			expected: map[int]bool{1: true, 2: true, 3: true, 5: true},
		},
	}

	for _, tc := range testTable {
		got := FindDifference(tc.setA, tc.setB)
		require.Equal(t, tc.expected, got)
	}
}
