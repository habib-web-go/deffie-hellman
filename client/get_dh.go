package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"time"

	pb "github.com/my/repo/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultNonce = "qwertyuiopasdfghjklz"
	defaultA     = "64"
	defaulta     = 6
	defaultP     = "1552518092300708935130918131258481755631334049434514313202351194902966239949102107258669453876591642442910007680288864229150803718918046342632727613031282983744380820890196288509170691316593175367469551763119843371637221007210577919"
	// defaultP = "23"
)

var (
	addr        = flag.String("addr", "localhost:5052", "the address to connect to")
	nonce       = flag.String("nonce", defaultNonce, "Nonce with length 20")
	serverNonce = flag.String("serverNonce", defaultNonce, "Server Nonce with length 20")
	A           = flag.String("A", defaultA, "A = g^a mod p")
	a           = flag.Uint("a", defaulta, "random number")
	p           = flag.String("p", defaultP, "Prime number")
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
	Abig, _ := new(big.Int).SetString(*A, 10)
	r, err := c.ReqDHParams(ctx, &pb.ReqDHParamsRequest{Nonce: *nonce, MessageId: 2, ServerNonce: *serverNonce, A: Abig.String()})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	log.Printf("Result: %s", r.String())
	p, _ := new(big.Int).SetString(*p, 10)
	B, _ := new(big.Int).SetString(r.GetB(), 10)
	log.Printf("sharedKey= %s",
		new(big.Int).Exp(
			B,
			big.NewInt(int64(*a)),
			p,
		).String(),
	)
}
