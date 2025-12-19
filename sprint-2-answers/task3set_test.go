package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombine(t *testing.T) {

	testTable := []struct {
		setA     map[string]bool
		setB     map[string]bool
		expected map[string]bool
	}{
		{
			setA: map[string]bool{
				"go": true, "java": true, "rust": true,
			},
			setB: map[string]bool{
				"rust": true, "python": true,
			},
			expected: map[string]bool{
				"go": true, "java": true, "rust": true, "python": true,
			},
		},
		{
			setA:     map[string]bool{},
			setB:     map[string]bool{},
			expected: map[string]bool{},
		},
		{
			setA: map[string]bool{
				"go": true, "java": true, "rust": true,
			},
			setB: map[string]bool{
				"go": true, "java": true, "rust": true,
			},
			expected: map[string]bool{
				"go": true, "java": true, "rust": true,
			},
		},
	}

	for _, tc := range testTable {
		got := Combine(tc.setA, tc.setB)
		require.Equal(t, tc.expected, got)
	}
}
