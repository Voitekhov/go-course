package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveElementsByConditions(t *testing.T) {

	testTable := []struct {
		ratings  map[int]int
		limit    int
		expected map[int]int
	}{
		{
			ratings: map[int]int{
				1: 10,
				2: 3,
				3: 7,
				4: 1,
			},
			limit: 5,
			expected: map[int]int{
				1: 10,
				3: 7,
			},
		},
		{
			ratings: map[int]int{
				1: 10,
				2: 3,
				3: 7,
				4: 1,
			},
			limit: 5,
			expected: map[int]int{
				1: 10,
				3: 7},
		},
		{
			ratings: map[int]int{
				1: 10,
				2: 3,
				3: 7,
				4: 1,
			},
			limit: 5,
			expected: map[int]int{
				1: 10,
				3: 7,
			},
		},
	}

	for _, tc := range testTable {
		RemoveElementsByConditions(tc.ratings, tc.limit)
		require.Equal(t, tc.expected, tc.ratings)
	}
}
