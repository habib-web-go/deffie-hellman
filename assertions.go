package main

import (
	"testing"
)

func AssertErrorMessagesEqual(
	t testing.TB,
	expected error,
	actual error,
) {
	if (expected == nil && actual != nil) || (expected != nil && actual == nil) {
		t.Errorf("Expected error message %v but got %v", expected, actual)
	} else if expected != nil && actual != nil {
		if expected.Error() != actual.Error() {
			t.Errorf(
				"Expected error message '%s' but got '%s'",
				expected.Error(),
				actual.Error(),
			)
		}
	}
}

func AssertStringLength(t testing.TB, str string, length int) {
	if len(str) != length {
		t.Errorf(
			"string length is %d, but expected length is %d",
			len(str),
			length,
		)
	}
}
