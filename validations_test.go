package main

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestCheckNonce(t *testing.T) {
	tests := []struct {
		name  string
		nonce string
		err   error
	}{
		{
			name:  "valid nonce",
			nonce: randomString(NonceLength),
			err:   nil,
		},
		{
			name:  "empty nonce",
			nonce: "",
			err: status.Errorf(
				codes.InvalidArgument,
				"Nonce length must be %d",
				NonceLength,
			),
		},
		{
			name:  "nonce too short",
			nonce: "abc",
			err: status.Errorf(
				codes.InvalidArgument,
				"Nonce length must be %d",
				NonceLength,
			),
		},
		{
			name:  "nonce too long",
			nonce: "abcdefghijklmnopqrstuvwxyz",
			err: status.Errorf(
				codes.InvalidArgument,
				"Nonce length must be %d",
				NonceLength,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkNonce(tt.nonce)
			AssertErrorMessagesEqual(t, err, tt.err)
		})
	}
}

func TestIsMessageIdValid(t *testing.T) {
	testCases := []struct {
		messageId uint64
		expected  bool
	}{
		{messageId: 4, expected: true},
		{messageId: 5, expected: false},
		{messageId: 0, expected: true},
	}

	for _, tc := range testCases {
		if output := isMessageIdValid(tc.messageId); output != tc.expected {
			t.Errorf(
				"Expected isMessageIdValid(%d) to be %v, got %v",
				tc.messageId,
				tc.expected,
				output,
			)
		}
	}
}
