package main

import (
	"log"
	"net"

	pb "demo/grpc"

	"github.com/pkg/errors"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) ProfileCreate(ctx context.Context, in *pb.ProfileCreateRequest) (*pb.ProfileCreateResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, errors.Wrap(err, "Cannot validate request while creating profile")
	}
	log.Printf("Hello %s %s", in.FirstName, in.LastName)
	return &pb.ProfileCreateResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting server on %s", port)
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
