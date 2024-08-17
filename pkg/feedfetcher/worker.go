package feedfetcher

import (
	"fmt"
	"sync"
	"time"

	"github.com/Zmahl/blog_aggregator/internal/database"
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
			go func(url string) {
				defer wg.Done()
				rss, err := FetchDataFromFeed(url)
				if err != nil {
					return
				}
				for _, item := range rss.Channel.Items {
					fmt.Println(item.Title)
				}
			}(feed.Url)
		}
	}
}
