proto: 
protoc --go_out=. --go-grpc_out=. auth/auth.prot

env:
	export PATH="$PATH:$(go env GOPATH)/bin"
