package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	pb "github.com/my/repo/grpc"

	"github.com/go-redis/redis"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

var (
	redisClient *redis.Client
)

func (s *server) reqPQ(ctx context.Context, req *pb.ReqPQRequest) (*pb.ReqPQResponse, error) {
	err := validateReqPQRequest(req)
	if err != nil {
		return nil, err
	}
	nonce := req.GetNonce()
	serverNonce := randomString(20)
	sha := createSHA1(nonce + serverNonce)
	p, g := initDeffieHellman()
	messageId := req.GetMessageId()
	jsonData, err := json.Marshal(clientHandShake{P: p, G: g, CurrentMessageId: messageId})
	if err != nil {
		log.Fatal("Failed to marshal struct to JSON:", err)
		return nil, err
	}
	setInRedis(sha, jsonData, time.Minute*20)

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
	a := req.GetA()
	b := randomUint()
	nonce := req.GetNonce()
	serverNonce := req.GetServerNonce()
	messageId := req.GetMessageId()

	handShakeData, err := getClientHandShake(nonce, serverNonce)
	B, sharedKey := createDeffieHellmanSharedKey(handShakeData.G, handShakeData.P, a, b)
	jsonData, err := json.Marshal(client{CurrentMessageId: messageId, AuthKey: sharedKey})
	if err != nil {
		log.Fatal("Failed to marshal struct to JSON:", err)
		return nil, err
	}
	setInRedis(strconv.Itoa(int(sharedKey)), jsonData, 0)
	return &pb.ReqDHParamsResponse{
		Nonce:       nonce,
		ServerNonce: serverNonce,
		MessageId:   messageId + 1,
		B:           B}, nil
}
func main() {
	loadEnv()
	runRedis()
	startServer()
}
