package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	bookspb "simple-grpc-client-server/books"
)

// This struct need to implement all RPC method of Service BookService
type BookServer struct {
	bookspb.UnimplementedBookServiceServer
}

func (b *BookServer) GetBooks(ctx context.Context, request *bookspb.BookRequest) (*bookspb.BookResponse, error) {
	response := bookspb.BookResponse{
		Name:      request.Name,
		Author:    "Yuval Harari",
		TotalPage: 450,
	}
	return &response, nil
}

func (b *BookServer) UpdateBooks(ctx context.Context, request *bookspb.BookRequest) (*bookspb.BookResponse, error) {
	log.Println("Your request to update books is accepted")
	updatedResponse := bookspb.BookResponse{
		Name:      request.Name,
		Author:    "Updated_Author",
		TotalPage: 100,
	}
	return &updatedResponse, nil
}

func main() {
	addr := "localhost:50051"
	// reserve the port.
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Println("Listening on  localhost:50051...")
	server := grpc.NewServer()
	bookspb.RegisterBookServiceServer(server, &BookServer{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
