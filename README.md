This is simple example how to create gRPC server and client. 

Use this command in current directory to generate the gRPC code:

```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative books/book.proto```

Start the book server:  
```go run books_server/main.go```

Call the book_server rpc method using client:  
```go run books_client/main.go```

