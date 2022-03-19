package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-go-hw/pkg/proto/hw"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMafiaClient(conn)

	var login string
	fmt.Printf("Введите логин: ")
	fmt.Scanf("%s", &login)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.AuthLogin(ctx, &pb.AuthReq{Name: login})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Успешная авторизация\n")

	r2, err := c.GetUsers(ctx, &pb.GetUsersReq{})
	fmt.Printf("Активные пользователи: %s\n", r2.GetUsers())
}
