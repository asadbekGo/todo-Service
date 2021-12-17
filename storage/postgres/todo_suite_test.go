package postgres

import (
	"testing"

	"github.com/asadbekGo/todo-service/config"
	pb "github.com/asadbekGo/todo-service/genproto"
	"github.com/asadbekGo/todo-service/pkg/db"
	"github.com/asadbekGo/todo-service/storage/repo"

	"github.com/stretchr/testify/suite"
)

type TodoRepositorySuite struct {
	suite.Suite
	CleanupFunc func()
	Repository  repo.TodoStorageI
}

func (suite *TodoRepositorySuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDBForSuite(config.Load())

	suite.Repository = NewTodoRepo(pgPool)
	suite.CleanupFunc = cleanup
}

// All methods that begin with "Test" are run as tests within a
// suite
func (suite *TodoRepositorySuite) TestTodoCRUD() {

	want := pb.Todo{
		Id:       "908b32e7-160f-4e6c-be3c-b1637a240b96",
		Assignee: "asadbek",
		Title:    "todo service",
		Summary:  "rpc implement",
		Deadline: "2021-12-15T14:12:14Z",
		Status:   "active",
	}

	todo := pb.Todo{
		Id:       "908b32e7-160f-4e6c-be3c-b1637a240b96",
		Assignee: "asadbek",
		Title:    "todo service",
		Summary:  "rpc implement",
		Deadline: "2021-12-15T14:12:14Z",
		Status:   "active",
	}

	_ = suite.Repository.Delete(want.Id)

	todo, err := suite.Repository.Create(todo)
	suite.Nil(err)

	getTodo, err := suite.Repository.Get(todo.Id)
	suite.Nil(err)
	suite.NotNil(getTodo, "getTodo must not be nil")
	suite.Equal(want.Id, getTodo.Id, "todo struct must match")

	want.Assignee = "Jack"

	updateTodo, err := suite.Repository.Update(want)
	suite.Nil(err)

	getTodo, err = suite.Repository.Get(want.Id)
	suite.Nil(err)
	suite.NotNil(getTodo)
	suite.Equal(want.Assignee, updateTodo.Assignee)

	listTodos, _, err := suite.Repository.List(1, 2)
	suite.Nil(err)
	suite.NotEmpty(listTodos)
	suite.Equal(want.Id, listTodos[0].Id)

	err = suite.Repository.Delete(want.Id)
	suite.Nil(err)
}

func (suite *TodoRepositorySuite) TearDownSuite() {
	suite.CleanupFunc()
}

func TestTodoRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TodoRepositorySuite))
}
