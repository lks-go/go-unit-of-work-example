package user

import (
	"context"
	"database/sql"
	"errors"

	"go-unit-of-work-example/internal/entity"
)

// QueryExecutor represents only used methods in this repository and implemented by sql.DB and sql.Tx
type QueryExecutor interface {
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func NewRepo(conn QueryExecutor) Repo {
	return Repo{conn: conn}
}

type Repo struct {
	conn QueryExecutor
}

func (r Repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := entity.User{}

	q := "SELECT id, name, email FROM users WHERE email = $1"
	err := r.conn.QueryRowContext(ctx, q, email).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrNotFound
		}

		return nil, err
	}

	return &u, nil
}

func (r Repo) AddNewUser(ctx context.Context, u entity.User) (string, error) {
	var id string

	q := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	err := r.conn.QueryRowContext(ctx, q, u.Name, u.Email).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
