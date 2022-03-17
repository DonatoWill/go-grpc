package main

import (
	"log"
	"context"
	"fmt"

	"github.com/donatowill/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id: "0",
		Name: "Joao",
		Email: "joao@example.com",
	}

	res, err := client.AddUser(context.Background(),req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}