package main

import (
	"context"
	"fmt"

	"github.com/alkmc/justforfunc/30_31_grpc/todo"

	"google.golang.org/protobuf/types/known/emptypb"
)

func add(ctx context.Context, client todo.TasksClient, text string) error {
	if _, err := client.Add(ctx, &todo.Text{Text: text}); err != nil {
		return fmt.Errorf("could not add task in the backend: %v", err)
	}

	fmt.Println("task added successfully")
	return nil
}

func list(ctx context.Context, client todo.TasksClient) error {
	l, err := client.List(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("could not fetch tasks: %v", err)
	}

	for _, t := range l.Tasks {
		if t.Done {
			fmt.Printf("👍")
		} else {
			fmt.Printf("😱")
		}
		fmt.Printf(" %s\n", t.Text)
	}

	return nil
}
