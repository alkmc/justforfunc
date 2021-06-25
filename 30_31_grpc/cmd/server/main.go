package main

import (
	"log"
	"net"

	"github.com/alkmc/justforfunc/30_31_grpc/todo"

	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	var tasks taskServer
	todo.RegisterTasksServer(srv, tasks)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Println("starting grpc service")
	log.Fatal(srv.Serve(l))
}
