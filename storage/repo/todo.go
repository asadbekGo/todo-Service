package repo

import (
	"time"

	pb "github.com/asadbekGo/todo-service/genproto"
)

// UserStorageI ...
type TodoStorageI interface {
	Create(pb.Todo) (pb.Todo, error)
	Get(id int64) (pb.Todo, error)
	List(page, limit int64) ([]*pb.Todo, int64, error)
	Update(pb.Todo) (pb.Todo, error)
	Delete(id int64) error
	ListOverdue(time time.Time, page, limit int64) ([]*pb.Todo, int64, error)
}
