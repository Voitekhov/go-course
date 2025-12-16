package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwapMap(t *testing.T) {

	testTable := []struct {
		numbers  map[string]int
		expected map[int]string
	}{
		{
			numbers:  map[string]int{},
			expected: map[int]string{},
		},
		{
			numbers: map[string]int{"user1": 10,
				"user2": 20,
				"user3": 30},
			expected: map[int]string{
				10: "user1",
				20: "user2",
				30: "user3",
			},
		},
		{
			numbers: map[string]int{
				"apple":  5,
				"banana": 1,
				"orange": 10},
			expected: map[int]string{
				5:  "apple",
				1:  "banana",
				10: "orange"},
		},
	}

	for _, tc := range testTable {
		got := SwapMap(tc.numbers)
		require.Equal(t, tc.expected, got)
	}
}
