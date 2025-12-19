package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildLetterMap(t *testing.T) {

	testTable := []struct {
		message  string
		expected map[rune]int
	}{
		{
			message: "hello world",
			expected: map[rune]int{
				'h': 1,
				'e': 1,
				'l': 3,
				'o': 2,
				'w': 1,
				'r': 1,
				'd': 1,
			},
		},
		{
			message:  "",
			expected: map[rune]int{},
		},
		{
			message: "aezakmi",
			expected: map[rune]int{
				'a': 2,
				'e': 1,
				'z': 1,
				'k': 1,
				'm': 1,
				'i': 1,
			},
		},
	}

	for _, tc := range testTable {
		got := BuildLetterMap(tc.message)
		require.Equal(t, tc.expected, got)
	}
}
