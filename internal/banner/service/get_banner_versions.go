package service

import (
	"context"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
)

func (s *BannerServiceApp) GetBannerVersions(ctx context.Context, ID uint64) ([]model.BannerVersion, error) {
	_, err := s.storage.GetBannerByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return s.storage.GetBannerVersions(ctx, ID)
}
