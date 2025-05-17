package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"go-unit-of-work-example/internal/entity"
)

type QueueProducer interface {
	Send(email, code string) error
}

type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	AddNewUser(ctx context.Context, u entity.User) (string, error)
}

type VerificationRepo interface {
	AddVerificationCode(ctx context.Context, userID string, code string) error
}

type RegisterUserUOW interface {
	InTransaction(ctx context.Context, f func(ctx context.Context, u UserRepo, v VerificationRepo) (err error)) error
}

func New(u UserRepo, r RegisterUserUOW, q QueueProducer) *Service {
	return &Service{
		userRepo: u,
		register: r,
		queue:    q,
	}
}

type Service struct {
	userRepo UserRepo
	register RegisterUserUOW
	queue    QueueProducer
}

func (s *Service) RegisterNewUser(ctx context.Context, name string, email string) error {
	u, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, entity.ErrNotFound) {
			return fmt.Errorf("failed to get user by email: %w", err)
		}
	}
	if u != nil {
		return entity.ErrAlreadyRegistered
	}

	err = s.register.InTransaction(ctx, func(ctx context.Context, u UserRepo, v VerificationRepo) error {
		id, err := u.AddNewUser(ctx, entity.User{Name: name, Email: email})
		if err != nil {
			return fmt.Errorf("failed to add new user: %w", err)
		}

		n, _ := rand.Int(rand.Reader, big.NewInt(1e6))
		code := fmt.Sprintf("%06s", n)

		if err := v.AddVerificationCode(ctx, id, code); err != nil {
			return fmt.Errorf("failed to add verification code: %w", err)
		}

		if err := s.queue.Send(email, code); err != nil {
			return fmt.Errorf("failed to send email: %w", err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
