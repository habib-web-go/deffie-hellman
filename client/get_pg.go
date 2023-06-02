package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"time"

	pb "github.com/my/repo/gen/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultNonce = "qwertyuiopasdfghjklz"
	defaultA     = 6
)

var (
	addr  = flag.String("addr", "localhost:5052", "the address to connect to")
	nonce = flag.String("nonce", defaultNonce, "Nonce with length 20")
	a     = flag.Int64("a", defaultA, "random number")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.ReqPQ(ctx, &pb.ReqPQRequest{Nonce: *nonce, MessageId: 0})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	log.Printf("Result: %s", r.String())
	p, _ := new(big.Int).SetString(r.GetP(), 10)
	log.Printf("A= %s", new(big.Int).Exp(
		big.NewInt(int64(r.GetG())),
		big.NewInt(*a),
		p))
}
