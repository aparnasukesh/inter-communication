proto: 
protoc --go_out=. --go-grpc_out=. auth/auth.proto

env:
	export PATH="$PATH:$(go env GOPATH)/bin"

