package main

import "testing"

func TestRandomString(t *testing.T) {
	testCases := []struct {
		name   string
		length int
	}{
		{
			name:   "valid input",
			length: 10,
		},
		{
			name:   "valid input",
			length: 5,
		},
		{
			name:   "empty input",
			length: 0,
		},
		{
			name:   "invalid input",
			length: 24,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := randomString(tc.length)
			AssertStringLength(t, str, tc.length)
		})
	}
}
