package main

import (
	"context"

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
	serverNonce := randomString(20)
	p, g := initDeffieHellman()

	// todo save in redis
	return &pb.ReqPQResponse{
		Nonce:       req.GetNonce(),
		ServerNonce: serverNonce,
		P:           p,
		G:           g,
		MessageId:   req.GetMessageId() + 1}, nil
}

func (s *server) ReqDHParams(ctx context.Context, req *pb.ReqDHParamsRequest) (*pb.ReqDHParamsResponse, error) {
	err := validateReqDHParamsRequest(req)
	if err != nil {
		return nil, err
	}
	a := req.GetA()
	b := randomUint()
	// todo reqrite in redis
	return &pb.ReqDHParamsResponse{
		Nonce:       req.GetNonce(),
		ServerNonce: req.GetServerNonce(),
		MessageId:   req.GetMessageId() + 1,
		B:           b}, nil
}
func main() {
	loadEnv()
	runRedis()
	startServer()
}
