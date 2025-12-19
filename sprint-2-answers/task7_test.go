package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvertBoolMap(t *testing.T) {

	testTable := []struct {
		mapToInvert map[string]bool
		expected    map[string]bool
	}{
		{
			mapToInvert: map[string]bool{
				"featureA": true,
				"featureB": false,
				"featureC": true,
			},
			expected: map[string]bool{
				"featureA": false,
				"featureB": true,
				"featureC": false,
			},
		},
		{
			mapToInvert: map[string]bool{},
			expected:    map[string]bool{},
		},
		{
			mapToInvert: map[string]bool{
				"Hello": true,
			},
			expected: map[string]bool{
				"Hello": false,
			},
		},
	}

	for _, tc := range testTable {
		got := InvertBoolMap(tc.mapToInvert)
		require.Equal(t, tc.expected, got)
	}
}
