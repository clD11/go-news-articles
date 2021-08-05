package storage

import (
	uuid "github.com/satori/go.uuid"
)

const (
	ErrInvalidKey = Error("db: invalid key")
	ErrNoRows     = Error("db: no rows in result set")
	ErrInsertRow  = Error("db: could not insert row")
)

type DB interface {
	Insert(key uuid.UUID, value interface{}) error
	Select(key uuid.UUID) (interface{}, error)
}

type SimpleDB struct {
	store map[uuid.UUID]interface{}
}

func NewDB() *SimpleDB {
	return &SimpleDB{
		store: make(map[uuid.UUID]interface{}),
	}
}

func (sdb *SimpleDB) Insert(key uuid.UUID, value interface{}) error {
	if len(key.Bytes()) <= 0 {
		return ErrInvalidKey
	}
	_, ok := sdb.store[key]
	if ok {
		return ErrInsertRow
	}
	sdb.store[key] = value
	return nil
}

func (sdb *SimpleDB) Select(key uuid.UUID) (interface{}, error) {
	if len(key.Bytes()) <= 0 {
		return nil, ErrInvalidKey
	}
	v, ok := sdb.store[key]
	if !ok {
		return nil, ErrNoRows
	}
	return v, nil
}

type Error string

func (e Error) Error() string {
	return string(e)
}
