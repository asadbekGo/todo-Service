package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pb "github.com/asadbekGo/todo-service/genproto"
)

type todoRepo struct {
	db *sqlx.DB
}

// NewTodoRepo ...
func NewTodoRepo(db *sqlx.DB) *todoRepo {
	return &todoRepo{db: db}
}

func (r *todoRepo) Create(todo pb.Todo) (pb.Todo, error) {
	var id int64
	err := r.db.QueryRow(`
		INSERT INTO todos(assignee, title, summary, deadline, status)
		VALUES ($1, $2, $3, $4, $5) returning id`,
		todo.Assignee,
		todo.Title,
		todo.Summary,
		todo.Deadline,
		todo.Status,
	).Scan(&id)
	if err != nil {
		return pb.Todo{}, err
	}

	todo, err = r.Get(id)

	if err != nil {
		return pb.Todo{}, nil
	}

	return todo, nil
}

func (r *todoRepo) Get(id int64) (pb.Todo, error) {
	var todo pb.Todo

	err := r.db.QueryRow(`
		SELECT id, assignee, title, summary, deadline, status FROM todos 
		WHERE id=$1`, id).Scan(
		&todo.Id,
		&todo.Assignee,
		&todo.Title,
		&todo.Summary,
		&todo.Deadline,
		&todo.Status,
	)
	if err != nil {
		return pb.Todo{}, err
	}

	return todo, nil
}

func (r *todoRepo) List(page, limit int64) ([]*pb.Todo, int64, error) {
	offset := (page - 1) * limit

	rows, err := r.db.Queryx(`
		SELECT id, assignee, title, summary, deadline, status FROM todos
		LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		todos []*pb.Todo
		count int64
	)

	for rows.Next() {
		var todo pb.Todo
		err = rows.Scan(
			&todo.Id,
			&todo.Assignee,
			&todo.Title,
			&todo.Summary,
			&todo.Deadline,
			&todo.Status,
		)
		if err != nil {
			return nil, 0, err
		}

		todos = append(todos, &todo)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM todos`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return todos, count, nil
}

func (r *todoRepo) Update(todo pb.Todo) (pb.Todo, error) {
	result, err := r.db.Exec(`
		UPDATE todos SET assignee=$1, title=$2, summary=$3, status=$4, deadline=$5
		WHERE id=$6`,
		todo.Assignee, todo.Title,
		todo.Summary, todo.Status,
		todo.Deadline, todo.Id)
	if err != nil {
		return pb.Todo{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Todo{}, sql.ErrNoRows
	}

	todo, err = r.Get(todo.Id)
	if err != nil {
		return pb.Todo{}, err
	}

	return todo, nil
}

func (r *todoRepo) Delete(id int64) error {
	result, err := r.db.Exec(`
		DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *todoRepo) ListOverdue(time time.Time, page, limit int64) ([]*pb.Todo, int64, error) {
	offset := (page - 1) * limit

	rows, err := r.db.Queryx(`
		SELECT id, assignee, title, summary, deadline, status FROM todos
		WHERE deadline > $1 LIMIT $2 OFFSET $3`, time, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		todos []*pb.Todo
		count int64
	)

	for rows.Next() {
		var todo pb.Todo

		err = rows.Scan(
			&todo.Id,
			&todo.Assignee,
			&todo.Title,
			&todo.Summary,
			&todo.Status,
			&todo.Deadline,
		)
		if err != nil {
			return nil, 0, err
		}

		todos = append(todos, &todo)
	}

	err = r.db.QueryRow(`
		SELECT count(*) FROM todos WHERE deadline > $1 LIMIT $2 OFFSET $3`,
		time, limit, offset).Scan(&count)

	if err != nil {
		return nil, 0, err
	}

	return todos, count, nil
}
