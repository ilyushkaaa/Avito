package service

import (
	"context"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
)

func (s *BannerServiceApp) GetUserBanner(ctx context.Context, featureID, tagID uint64, lastVersion bool) (*model.Banner, error) {
	banner, err := s.storage.GetBannerByFeatureTag(ctx, featureID, tagID)
	if err != nil {
		return nil, err
	}
	if !banner.IsActive {
		return nil, ErrBannerIsInactive
	}

	feature, tags, err := s.storage.GetBannerFeatureTags(ctx, banner.ID)
	if err != nil {
		return nil, err
	}

	banner.FeatureID = feature
	banner.TagIDs = tags

	return banner, nil
}
