package entity

import "errors"

var (
	ErrNotFound          = errors.New("not found")
	ErrAlreadyRegistered = errors.New("user already registered")
)
