package sprint_2_answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoginCheck(t *testing.T) {

	testTable := []struct {
		login    string
		users    map[string]bool
		expected bool
	}{
		{
			users:    map[string]bool{"Alice": true, "Bob": false},
			login:    "Alice",
			expected: true,
		},
		{
			users:    map[string]bool{"Alice": true, "Bob": false},
			login:    "Bob",
			expected: true,
		},
		{
			users:    map[string]bool{"Alice": true, "Bob": false},
			login:    "Charlie",
			expected: false,
		},
	}

	for _, tc := range testTable {
		got := LoginCheck(tc.users, tc.login)
		require.Equal(t, tc.expected, got)
	}
}
