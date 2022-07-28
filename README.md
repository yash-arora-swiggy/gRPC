# gRPC

If creating a new repo gRPC directory use - go mod init github.com/<username>/package
Use go get ./...
To autogenerate the go code - cd to calculator directory and run protoc --go_out=. --go-grpc_out=. ./calculator.proto


First Run the server from server directory - go run server.go
To run client go to client directory - go run client.go
  
