gen:
	protoc --go_out=. --go-grpc_out=. pb/sum.proto

rs:
	go run server/server.go

rc:
	go run client/client.go