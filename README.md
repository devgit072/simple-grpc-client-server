This is simple example how to create gRPC server and client. 

Use this command in current directory to generate the gRPC code:

```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative books/book.proto```

Start the book server:  
```go run books_server/main.go```

Call the book_server rpc method using client:  
```go run books_client/main.go```

<h3> Streaming API: </h3>
There are four types api supported in gRPC:

1) <b>Unary api</b>: Client sends one request and Server responds back with one response in one TCP connecttion.
2) <b>Server streaming api</b>: Client sends one request and server responds back with stream of response.
3) <b>Client streaming api</b>: Client sends streams of data as request and server responds back with one response over one TCP connection.
4) <b>Bi-Directional streaming api</b>: Both client and server sends stream as request and response respectively.

Examples of streaming api implementaion can be found in <b>streaming_api</b> folder.

Generate the grpc code in streaming_folder: 

```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative streaming_api/protos/*.proto```
 
Client code can be found in <b>streaming_api/client</b> and server code can be found in <b>streaming_api/server</b>






