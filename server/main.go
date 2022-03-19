package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	pb "grpc-go-hw/pkg/proto/hw"
	"log"
	"net"
	"net/http"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

var (
	users []string
)

type server struct {
	pb.UnimplementedMafiaServer
}

func (s *server) AuthLogin(ctx context.Context, in *pb.AuthReq) (*pb.AuthResp, error) {
	login := in.GetName()
	log.Printf("Received: %s", login)

	for _, el := range users {
		if el == login {
			log.Printf("Login %s is busy", login)
			return nil, status.Error(http.StatusUnauthorized, "Login is busy")
		}
	}

	users = append(users, login)
	return &pb.AuthResp{}, nil
}

func (s *server) GetUsers(ctx context.Context, in *pb.GetUsersReq) (*pb.GetUsersResp, error) {
	return &pb.GetUsersResp{Users: users}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMafiaServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
