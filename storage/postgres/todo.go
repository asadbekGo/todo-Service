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
	var id string
	err := r.db.QueryRow(`
		INSERT INTO todos(id, assignee, title, summary, deadline, status, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) returning id`,
		todo.Id,
		todo.Assignee,
		todo.Title,
		todo.Summary,
		todo.Deadline,
		todo.Status,
		time.Now().UTC(),
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

func (r *todoRepo) Get(id string) (pb.Todo, error) {
	var todo pb.Todo
	err := r.db.QueryRow(`
		SELECT id, assignee, title, summary, deadline, status, created_at, updated_at
		FROM todos 
		WHERE id=$1 and deleted_at is null`, id).Scan(
		&todo.Id,
		&todo.Assignee,
		&todo.Title,
		&todo.Summary,
		&todo.Deadline,
		&todo.Status,
		&todo.CreatedAt,
		&todo.UpdatedAt,
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
		WHERE deleted_at is null
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

	err = r.db.QueryRow(`SELECT count(*) FROM todos WHERE deleted_at is null`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return todos, count, nil
}

func (r *todoRepo) Update(todo pb.Todo) (pb.Todo, error) {
	result, err := r.db.Exec(`
		UPDATE todos SET assignee=$1, title=$2, summary=$3, status=$4, deadline=$5, updated_at=$6
		WHERE id=$7 and deleted_at is null`,
		todo.Assignee, todo.Title,
		todo.Summary, todo.Status,
		todo.Deadline, time.Now().UTC(), todo.Id)
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

func (r *todoRepo) Delete(id string) error {
	result, err := r.db.Exec(`
		UPDATE todos SET deleted_at=current_timestamp WHERE id=$1`, id)
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
		WHERE deadline > $1 and deleted_at is null LIMIT $2 OFFSET $3`, time, limit, offset)
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

	err = r.db.QueryRow(`
		SELECT count(*) FROM todos WHERE deadline > $1 and deleted_at is null LIMIT $2 OFFSET $3`,
		time, limit, offset).Scan(&count)

	if err != nil {
		return nil, 0, err
	}

	return todos, count, nil
}
