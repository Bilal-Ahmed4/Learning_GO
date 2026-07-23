package repository

import (
	"context"
	"time"

	"github.com/Bilal-Ahmed4/to-do-app/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodo(pool *pgxpool.Pool, title string, completed bool) (*models.Todo, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `INSERT INTO todos (title,completed)
						VALUES ($1, $2)
						RETURNING id ,title ,completed,created_at,updated_at;`

	var todo models.Todo
	err := pool.QueryRow(ctx, query, title, completed).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil

}

func GetTodos(pool *pgxpool.Pool) ([]*models.Todo, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	var query string = `SELECT * FROM todos;`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var todos []*models.Todo

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func GetTodoById(pool *pgxpool.Pool, id int) (*models.Todo, error) {
	ctx := context.Background()
	ctx, cancle := context.WithTimeout(ctx, 5*time.Second)
	defer cancle()

	var query string = `SELECT * FROM todos
					  WHERE id=$1;`

	var todo models.Todo
	err := pool.QueryRow(ctx, query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil

}
