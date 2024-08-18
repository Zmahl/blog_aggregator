package feedfetcher

import (
	"encoding/xml"
	"net/http"
	"time"
)

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	PubDate     time.Time `xml:"pubdate"`
}

func FetchDataFromFeed(url string) (RSS, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RSS{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return RSS{}, err
	}
	defer resp.Body.Close()

	rss := RSS{}
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)

	if err != nil {
		return RSS{}, err
	}

	return rss, nil
}
