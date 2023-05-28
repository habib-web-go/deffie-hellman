package main

import (
	"log"
	"math/rand"
	"net"
	"testing"
	"time"

	pb "github.com/my/repo/grpc"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func AssertErrorMessagesEqual(
	t testing.TB,
	expected error,
	actual error,
) {
	if (expected == nil && actual != nil) || (expected != nil && actual == nil) {
		t.Errorf("Expected error message %v but got %v", expected, actual)
	} else if expected != nil && actual != nil {
		if expected.Error() != actual.Error() {
			t.Errorf("Expected error message '%s' but got '%s'", expected.Error(), actual.Error())
		}
	}
}

func RandomString(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	const charset = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}

	return string(b)
}

func startServer() {

	lis, err := net.Listen("tcp", ":5052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})

	log.Println("gRPC server listening on port 5052")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
