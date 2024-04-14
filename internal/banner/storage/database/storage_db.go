package storage

import (
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/database/postgres/database"
	"go.uber.org/zap"
)

type BannerStorageDB struct {
	db     database.Database
	logger *zap.SugaredLogger
}

func New(db database.Database, logger *zap.SugaredLogger) *BannerStorageDB {
	return &BannerStorageDB{db: db, logger: logger}
}
