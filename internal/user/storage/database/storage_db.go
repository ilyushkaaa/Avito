package storage

import (
	"github.com/ilyushkaaa/banner-service/internal/database/postgres/database"
)

type UserStorageDB struct {
	db database.Database
}

func New(db database.Database) *UserStorageDB {
	return &UserStorageDB{db: db}
}
