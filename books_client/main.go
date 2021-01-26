package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	bookspb "simple-grpc-client-server/books"
)

func main() {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		return
	}
	defer conn.Close()

	client := bookspb.NewBookServiceClient(conn)
	req := bookspb.BookRequest{Name: "Sapiens"}
	resp, err := client.GetBooks(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	log.Println("Response: ", resp)
}
