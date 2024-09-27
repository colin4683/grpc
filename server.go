package main

import (
	"context"
	"log"
	"net"

	pb "github.com/colin4683/grpc/protos"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Authenticate(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Printf("Received: %v", in.GetUsername())
	return &pb.AuthResponse{Token: "token"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
