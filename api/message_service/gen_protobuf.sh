protoc -I. --go_out=../../pkg/message_service --go_opt=paths=source_relative --go-grpc_out=../../pkg/message_service --go-grpc_opt=paths=source_relative api.proto
