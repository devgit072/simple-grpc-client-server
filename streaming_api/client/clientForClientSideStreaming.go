package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	employeepb "simple-grpc-client-server/streaming_api/protos"
	"time"
)

func clientStreamingDemo() {
	log.Println("Client for ClientSide streaming...")
	log.Println("Client for server streaming")

	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(addr, opts)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer conn.Close()
	client := employeepb.NewEmployeeServiceClient(conn)

	// First get the stream object.
	stream, err := client.ClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	reqs := []*employeepb.Employee{
		{Name: "Devraj-1", Age: 20, Address: "Blr-1"},
		{Name: "Devraj-2", Age: 21, Address: "Blr-2"},
		{Name: "Devraj-3", Age: 22, Address: "Blr-3"},
		{Name: "Devraj-4", Age: 23, Address: "Blr-4"},
		{Name: "Devraj-5", Age: 24, Address: "Blr-5"},
		{Name: "Devraj-6", Age: 25, Address: "Blr-6"},
		{Name: "Devraj-7", Age: 26, Address: "Blr-7"},
	}

	for _, r := range reqs {
		time.Sleep(2 * time.Second)
		if err := stream.Send(r); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
	}
	// Now get one response from server.
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Printf("Response from server: %t\n", resp.Success)
}
