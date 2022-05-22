package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	v1 "github.com/fumist23/test-envoy/api/v1"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port.")
)

type server struct {
	v1.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloResponse, error) {
	log.Printf("name: %s", req.GetName())
	return &v1.HelloResponse{
		Message: fmt.Sprintf("Hello %s!!!", req.GetName()),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	v1.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
