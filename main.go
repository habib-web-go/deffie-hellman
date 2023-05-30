package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/my/repo/grpc"

	"github.com/go-redis/redis"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

var (
	redisClient *redis.Client
)

func (s *server) ReqPQ(ctx context.Context, req *pb.ReqPQRequest) (*pb.ReqPQResponse, error) {
	err := validateReqPQRequest(req)
	if err != nil {
		return nil, err
	}
	nonce := req.GetNonce()
	serverNonce := randomString(20)
	sha := createSHA1(nonce + serverNonce)
	p, g, err := initDeffieHellman()
	if err != nil {
		return nil, err
	}
	messageId := req.GetMessageId()
	jsonData, err := json.Marshal(clientHandShake{P: p, G: g, CurrentMessageId: messageId})
	if err != nil {
		log.Fatal("Failed to marshal struct to JSON:", err)
		return nil, err
	}
	setInRedis(sha, jsonData, ExpirationDuration)
	return &pb.ReqPQResponse{
		Nonce:       nonce,
		ServerNonce: serverNonce,
		P:           p,
		G:           g,
		MessageId:   messageId + 1}, nil
}

func (s *server) ReqDHParams(ctx context.Context, req *pb.ReqDHParamsRequest) (*pb.ReqDHParamsResponse, error) {
	err := validateReqDHParamsRequest(req)
	if err != nil {
		return nil, err
	}
	A := req.GetA()
	b := randomUint()
	nonce := req.GetNonce()
	serverNonce := req.GetServerNonce()
	messageId := req.GetMessageId()

	handShakeData, err := getClientHandShake(nonce, serverNonce)
	B, sharedKey := createDeffieHellmanSharedKey(handShakeData.G, handShakeData.P, A, b)
	jsonData, err := json.Marshal(client{CurrentMessageId: messageId, AuthKey: sharedKey})
	if err != nil {
		log.Fatal("Failed to marshal struct to JSON:", err)
		return nil, err
	}
	setInRedis(sharedKey, jsonData, ExpirationDuration)
	return &pb.ReqDHParamsResponse{
		Nonce:       nonce,
		ServerNonce: serverNonce,
		MessageId:   messageId + 1,
		B:           B}, nil
}

func (s *server) IsValidAuthkey(ctx context.Context, req *pb.IsValidAuthKeyRequest) (*pb.IsValidAuthKeyResponse, error) {
	isValid, err := existInRedis(req.GetAuthkey())
	if err != nil {
		log.Fatal("Failed to connect to redis:", err)
		return nil, err
	}
	return &pb.IsValidAuthKeyResponse{IsValid: isValid}, nil
}

func main() {
	loadEnv()
	runRedis()
	startServer()
}
