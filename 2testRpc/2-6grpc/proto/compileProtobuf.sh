#protoc --go-grpc_out=:. --go_out=. ./user.proto
protoc --go-grpc_out=require_unimplemented_servers=false:. --go_out=. ./user.proto