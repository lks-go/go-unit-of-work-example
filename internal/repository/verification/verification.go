package verification

import (
	"context"
	"database/sql"
)

// QueryExecutor represents only used methods in this repository and implemented by sql.DB and sql.Tx
type QueryExecutor interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

func NewRepo(conn QueryExecutor) Repo {
	return Repo{conn: conn}
}

type Repo struct {
	conn QueryExecutor
}

func (s Repo) AddVerificationCode(ctx context.Context, userID string, code string) error {
	q := "INSERT INTO verification (user_id, code, vefified) VALUES ($1, $2)"
	if _, err := s.conn.ExecContext(ctx, q, userID, code, false); err != nil {
		return err
	}

	return nil
}
