package db

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrAlreadyExist = errors.New("already exist")
)

type DB interface {
	Get(id string) (string, error)
	Set(id, name string) error
}
