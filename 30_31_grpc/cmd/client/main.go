package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/alkmc/justforfunc/30_31_grpc/todo"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("missing subcommand: list or add")
	}

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to backend: %v", err)
	}
	client := todo.NewTasksClient(conn)

	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list(context.Background(), client)
	case "add":
		err = add(context.Background(), client, strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}
	if err != nil {
		log.Fatal(err)
	}
}
