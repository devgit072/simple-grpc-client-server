package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	employeepb "simple-grpc-client-server/streaming_api/protos"
)

// Creating Client for server streaming response
func serverStreamingDemo() {
	log.Println("Client for server streaming")

	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(addr, opts)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer conn.Close()
	client := employeepb.NewEmployeeServiceClient(conn)
	req := employeepb.EmployeesRequest{NumberOfEmployees: 10}
	serverResponseStream, err := client.ServerSreaming(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	for {
		msg, err := serverResponseStream.Recv()
		if err == io.EOF {
			log.Println("Stream response is over")
			break
		}
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		log.Printf("Employee Response: %+v\n", msg)
	}
}
