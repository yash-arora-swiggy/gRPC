# gRPC


## To install go dependencies for protoc
1. `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
2. `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
3. `export PATH=$PATH:~/go/bin`

## If creating a new repo gRPC directory use - 
1. `go mod init github.com/<username>/package`
2. `go get ./...`

## To autogenerate the go code - 
cd to calculator directory and run `protoc --go_out=. --go-grpc_out=. ./calculator.proto`


###### First Run the server from server directory - `go run server.go`
###### To run client go to client directory - `go run client.go`
  
