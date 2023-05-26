package main

import (
	"context"

	pb "grpc/auth/deffieHellman"

	"github.com/go-redis/redis"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

var (
	redisClient *redis.Client
)

func (s *server) SayHello(ctx context.Context, req *pb.ReqPQRequest) (*pb.ReqPQResponce, error) {
	name := req.GetName()
	message := "Hello, " + name
	return &pb.ReqPQResponse{Message: message}, nil
}

func main() {
	loadEnv()
	runRedis()
	startServer()
}
