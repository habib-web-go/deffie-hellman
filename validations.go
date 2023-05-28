package main

import (
	pb "github.com/my/repo/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkNonce(nonce string) error {
	if !hasNonceValidLength(nonce) {
		return status.Errorf(
			codes.InvalidArgument,
			"Nonce length must be %d",
			NonceLength,
		)
	}
	return nil
}

func hasNonceValidLength(nonce string) bool {
	return len(nonce) == NonceLength
}

func checkClientMessageId(messageId uint64, nonce string, serverNonce string) error {
	if !isMessageIdValid(messageId) {
		return status.Errorf(
			codes.InvalidArgument,
			"MessageId must be even",
		)
	}

	if messageId == 0 {
		return nil
	}

	// TODO: If is in getting auth key check message id with nonce or check auth key?!

	return nil
}

func isMessageIdValid(messageId uint64) bool {
	return messageId%2 == 0
}

func validateReqPQRequest(req *pb.ReqPQRequest) error {
	err := checkNonce(req.GetNonce())
	if err != nil {
		return err
	}

	err = checkClientMessageId(
		req.GetMessageId(),
		req.GetNonce(),
		"",
	)
	if err != nil {
		return err
	}

	return nil
}

func validateReqDHParamsRequest(req *pb.ReqDHParamsRequest) error {
	err := checkNonce(req.GetNonce())
	if err != nil {
		return err
	}

	err = checkNonce(req.GetServerNonce())
	if err != nil {
		return err
	}

	err = checkClientMessageId(
		req.GetMessageId(),
		req.GetNonce(),
		req.GetServerNonce(),
	)
	if err != nil {
		return err
	}

	return nil
}
