package feedfetcher

import (
	"context"

	"github.com/Zmahl/blog_aggregator/internal/database"
)

func GetNextFeedsToFetch(feedAmount int, db *database.Queries) ([]database.Feed, error) {
	feeds, err := db.GetNextFeedsFetch(context.Background(), int32(feedAmount))
	if err != nil {
		return nil, err
	}
	return feeds, nil
}
