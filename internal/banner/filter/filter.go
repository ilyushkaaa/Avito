package filter

type Filter struct {
	FeatureID uint64
	TagID     uint64
	Offset    uint64
	Limit     uint64
}

func New(featureID, tagId, offset, limit uint64) Filter {
	return Filter{
		FeatureID: featureID,
		TagID:     tagId,
		Offset:    offset,
		Limit:     limit,
	}
}
