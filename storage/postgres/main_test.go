package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/asadbekGo/todo-service/config"
	"github.com/asadbekGo/todo-service/pkg/db"
	"github.com/asadbekGo/todo-service/pkg/logger"
)

var pgRepo *todoRepo

func TestMain(m *testing.M) {
	cfg := config.Load()

	connDB, err := db.ConnectionToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	pgRepo = NewTodoRepo(connDB)

	os.Exit(m.Run())
}
