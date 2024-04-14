package service

import (
	"context"
)

func (s *BannerServiceApp) DeleteBanner(ctx context.Context, ID uint64) error {
	return s.storage.DeleteBanner(ctx, ID)
}
