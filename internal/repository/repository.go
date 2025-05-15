package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-unit-of-work-example/internal/repository/user"
	"go-unit-of-work-example/internal/repository/verification"
	"go-unit-of-work-example/internal/service"
)

func NewRegisterUOW(conn *sql.DB) *RegisterUOW {
	return &RegisterUOW{conn: conn}
}

type RegisterUOW struct {
	conn *sql.DB
}

func (r RegisterUOW) InTransaction(ctx context.Context, fn func(ctx context.Context, u service.UserRepo, v service.VerificationRepo) error) error {
	tx, err := r.conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				// TODO use your favorite logger
				fmt.Printf("failed to rollback transaction: %s", err)
			}
			return
		}

		if err := tx.Commit(); err != nil {
			// TODO use your favorite logger
			fmt.Printf("failed to commit transaction: %s", err)
		}
	}()

	if err := fn(ctx, user.NewRepo(tx), verification.NewRepo(tx)); err != nil {
		return fmt.Errorf("failed to execute transaction: %w", err)
	}

	return nil
}
