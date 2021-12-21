package service

import (
	"log"
	"os"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/asadbekGo/todo-service/genproto"
)

var client pb.TodoServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	client = pb.NewTodoServiceClient(conn)

	os.Exit(m.Run())
}
