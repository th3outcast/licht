package errors

import "errors"

var (
	ErrNotFound = errors.New("Not found")
	ErrSetKV = errors.New("Error setting key-value")
)
