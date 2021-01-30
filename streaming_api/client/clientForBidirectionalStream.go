package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	employeepb "simple-grpc-client-server/streaming_api/protos"
	"time"
)

func bidirectionalStreamingDemo() {
	reqs := []*employeepb.Employee{
		{Name: "Devraj-1", Age: 20, Address: "Blr-1"},
		{Name: "Devraj-2", Age: 21, Address: "Blr-2"},
		{Name: "Devraj-3", Age: 22, Address: "Blr-3"},
		{Name: "Devraj-4", Age: 23, Address: "Blr-4"},
		{Name: "Devraj-5", Age: 24, Address: "Blr-5"},
		{Name: "Devraj-6", Age: 25, Address: "Blr-6"},
		{Name: "Devraj-7", Age: 26, Address: "Blr-7"},
	}

	log.Println("Client for ClientSide streaming...")
	log.Println("Client for server streaming")

	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(addr, opts)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer conn.Close()
	client := employeepb.NewEmployeeServiceClient(conn)
	stream, err := client.BidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	ch := make(chan struct{})
	// Send requests in goroutine.
	go func() {
		for _, req := range reqs {
			fmt.Printf("Sending request to update employee: %+v\n", req)
			if err := stream.Send(req); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			time.Sleep(2 * time.Second)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
	}()

	// Now get the stream response from server.
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				close(ch)
				log.Fatalf("Error: %s", err.Error())
			}
			log.Printf("Response from server: %+v\n", res)
		}
		close(ch)
	}()
	<-ch
}
