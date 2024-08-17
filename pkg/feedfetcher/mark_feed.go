package feedfetcher

import (
	"context"
	"time"

	"github.com/Zmahl/blog_aggregator/internal/database"
)

func MarkFeedFetched(feeds []database.Feed, db *database.Queries) {
	now := time.Now().UTC()
	ctx := context.Background()
	for _, f := range feeds {
		db.UpdateFeed(ctx, database.UpdateFeedParams{
			ID:        f.ID,
			UpdatedAt: now,
		})
	}
}
