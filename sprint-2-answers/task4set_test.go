package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckSubSet(t *testing.T) {

	testTable := []struct {
		setA     map[int]bool
		setB     map[int]bool
		expected bool
	}{
		{
			setA: map[int]bool{
				1: true, 2: true,
			},
			setB: map[int]bool{
				1: true, 2: true, 3: true, 4: true,
			},
			expected: true,
		},
		{
			setA:     map[int]bool{},
			setB:     map[int]bool{},
			expected: true,
		},
		{
			setA: map[int]bool{
				1: true, 2: true, 3: true, 4: true,
			},
			setB: map[int]bool{
				1: true, 2: true,
			},
			expected: false,
		},
	}

	for _, tc := range testTable {
		got := CheckSubSet(tc.setA, tc.setB)
		require.Equal(t, tc.expected, got)
	}
}
