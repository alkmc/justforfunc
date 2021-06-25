# to generate proto

protoc --go_out=. --go-grpc_out=. todo.proto --go-grpc_opt=require_unimplemented_servers=false