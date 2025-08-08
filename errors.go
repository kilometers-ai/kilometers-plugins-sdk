package kmsdk

import "errors"

var (
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInvalidConfig  = errors.New("invalid config")
	ErrNotInitialized = errors.New("plugin not initialized")
)
