package feedfetcher

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Zmahl/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

type Worker struct {
	Interval    int
	NumberFeeds int
}

func (w *Worker) FetchAndUpdateFeeds(db *database.Queries) {
	for {
		time.Sleep(time.Duration(w.Interval) * time.Second)
		feeds, err := GetNextFeedsToFetch(w.NumberFeeds, db)
		if err != nil {
			return
		}

		MarkFeedFetched(feeds, db)
		wg := sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			ctx := context.Background()
			go func(f database.Feed) {
				defer wg.Done()
				rss, err := FetchDataFromFeed(f.Url)
				if err != nil {
					return
				}
				for _, item := range rss.Channel.Items {
					_, err := db.CreatePost(ctx, database.CreatePostParams{
						ID:          uuid.New(),
						CreatedAt:   time.Now().UTC(),
						UpdatedAt:   time.Now().UTC(),
						Title:       item.Title,
						Description: item.Description,
						PublishedAt: item.PubDate,
						FeedID:      f.ID,
					})
					if err != nil {
						log.Println(err.Error())
					}
				}
			}(feed)
		}
	}
}
