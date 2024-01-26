package services

import (
	"backend_01/food/proto"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

var client proto.UserServiceClient

func GetFoods() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewUserServiceClient(conn)

	// contact the server and print out its response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := client.GetAllUsers(ctx, &proto.FoodRequest{})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Default().Println(r)
}
