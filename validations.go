package main

import (
	pb "github.com/my/repo/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkNonce(nonce string) error {
	isValid := len(nonce) == NONCE_LENGTH
	if !isValid {
		return status.Errorf(codes.InvalidArgument, "Nonce length must be %d", NONCE_LENGTH)
	}
	return nil
}

func checkClientMessageId(messageId uint64, nonce string, serverNonce string) error {
	isValid := messageId%2 == 0
	if !isValid {
		return status.Errorf(codes.InvalidArgument, "MessageId must be even")
	}

	if messageId == 0 {
		return nil
	}

	clientData, err := getClientData(nonce, serverNonce)
	if err != nil {
		return err
	}
	if clientData.CurrentMessageId >= uint64(messageId) {
		return status.Errorf(codes.InvalidArgument, "Not valid messageId")
	}

	return nil
}

func validateReqPQRequest(req *pb.ReqPQRequest) error {
	err := checkNonce(req.GetNonce())
	if err != nil {
		return err
	}
	err = checkClientMessageId(req.GetMessageId(), req.GetNonce(), "")
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
	err = checkClientMessageId(req.GetMessageId(), req.GetNonce(), req.GetServerNonce())
	if err != nil {
		return err
	}
	return nil
}
