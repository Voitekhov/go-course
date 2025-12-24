package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersect(t *testing.T) {

	testTable := []struct {
		bigSet   map[int]bool
		smallSet map[int]bool
		expected map[int]bool
	}{
		{
			bigSet: map[int]bool{
				1: true, 2: true, 3: true, 5: true,
			},
			smallSet: map[int]bool{
				3: true, 4: true, 5: true, 6: true,
			},
			expected: map[int]bool{
				3: true, 5: true,
			},
		},
		{
			bigSet:   map[int]bool{},
			smallSet: map[int]bool{},
			expected: map[int]bool{},
		},
		{
			bigSet: map[int]bool{
				1: true, 2: true, 3: true, 5: true,
			},
			smallSet: map[int]bool{
				10: true, 4: true, 11: true, 6: true,
			},
			expected: map[int]bool{},
		},
		{
			bigSet: map[int]bool{},
			smallSet: map[int]bool{
				3: true, 4: true, 5: true, 6: true,
			},
			expected: map[int]bool{},
		},
		{
			bigSet: map[int]bool{
				1: true, 2: true, 3: true, 5: true,
			},
			smallSet: map[int]bool{},
			expected: map[int]bool{},
		},
	}

	for _, tc := range testTable {
		got := Intersect(tc.bigSet, tc.smallSet)
		require.Equal(t, tc.expected, got)
	}
}
