gen-protoc:
	protoc --go_out=. --go-grpc_out=. ./proto/users/user.proto

swag-gen:
  ~/go/bin/swag init -g ./api/router/router.go -o api/swagger/docs force 1
