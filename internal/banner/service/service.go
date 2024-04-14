package service

import (
	"context"

	"github.com/ilyushkaaa/banner-service/internal/banner/filter"
	"github.com/ilyushkaaa/banner-service/internal/banner/model"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
)

type BannerService interface {
	AddBanner(ctx context.Context, banner model.Banner) (*model.Banner, error)
	GetBanners(ctx context.Context, filter filter.Filter) ([]model.Banner, error)
	GetUserBanner(ctx context.Context, featureID, tagID uint64, lastVersion bool) (*model.Banner, error)
	UpdateBanner(ctx context.Context, banner BannerToUpdate) error
	DeleteBanner(ctx context.Context, ID uint64) error
	GetBannerVersions(ctx context.Context, ID uint64) ([]model.BannerVersion, error)
	ApplyBannerVersion(ctx context.Context, versionID uint64) error
}

type BannerServiceApp struct {
	storage storage.BannerStorage
}

func New(storage storage.BannerStorage) *BannerServiceApp {
	return &BannerServiceApp{storage: storage}
}
