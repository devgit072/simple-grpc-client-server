package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	employeepb "simple-grpc-client-server/streaming_api/protos"
	"time"
)

const addr = "localhost:50051"

type server struct {
	employeepb.UnimplementedEmployeeServiceServer
}

// Implement rpc which will send stream response.
func (srv *server) ServerSreaming(req *employeepb.EmployeesRequest, stream employeepb.EmployeeService_ServerSreamingServer) error {
	log.Println("Server streaming RPC in server...")
	totalReq := int(req.NumberOfEmployees)

	for i := 0; i < totalReq; i++ {
		// Just generating some random response for demo purpose.
		resp := employeepb.EmployeeResponse{
			Id:      int64(i + 1),
			Name:    fmt.Sprintf("Employee-%d", i),
			Age:     int32(20 + i),
			Address: fmt.Sprintf("Address- %d", i),
		}
		log.Println("Sending one response from server...")
		if err := stream.Send(&resp); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		log.Println("Sent one response")
		time.Sleep(2 * time.Second)
	}
	return nil
}

func (srv *server) ClientStreaming(stream employeepb.EmployeeService_ClientStreamingServer) error {
	// accepts the stream requests came from server.
	log.Println("Server-side RPC accepting client streaming request...")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&employeepb.Response{Success: true})
		}
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		log.Printf("Updated the employee: %+v", req)
	}
}

func (srv *server) BidirectionalStreaming(stream employeepb.EmployeeService_BidirectionalStreamingServer) error {
	log.Println("Server-side RPC for bidirectional streaming...")
	// Accepts the stream requests from Client.
	i := 1 // Just random employee id
	for {
		req, err := stream.Recv()
		// Always check for EOF.
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		i++
		updatedEmployee := employeepb.EmployeeResponse{
			Id:      int64(i),
			Name:    req.Name,
			Age:     req.Age,
			Address: req.Address,
		}
		log.Printf("Updated the employee with details: %+v", updatedEmployee)
		if err := stream.Send(&updatedEmployee); err != nil {
			log.Fatalf("Error: %s", err.Error())
			return err
		}
	}
}

// starting the server
func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Println("Listening on ", addr, ".....")
	s := grpc.NewServer()
	employeepb.RegisterEmployeeServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
