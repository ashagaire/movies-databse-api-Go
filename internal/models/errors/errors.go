package models

import "errors"

var (
	ErrNotFound      = errors.New("resource not found")
	ErrInvalidInput  = errors.New("invalid input provided")
	ErrConflict      = errors.New("resource has existing relationships")
)
