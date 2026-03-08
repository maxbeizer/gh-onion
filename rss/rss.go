package rss

import (
	"encoding/xml"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

const defaultFeedURL = "https://www.theonion.com/rss"

// Item represents a single RSS feed item.
type Item struct {
	Title   string `xml:"title" json:"title"`
	Link    string `xml:"link" json:"link"`
	PubDate string `xml:"pubDate" json:"pub_date,omitempty"`
}

type rssChannel struct {
	Items []Item `xml:"item"`
}

type rssFeed struct {
	Channel rssChannel `xml:"channel"`
}

// Fetch retrieves and parses headlines from The Onion's RSS feed.
func Fetch() ([]Item, error) {
	return FetchURL(defaultFeedURL)
}

// FetchURL retrieves and parses headlines from a given RSS URL.
func FetchURL(url string) ([]Item, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching RSS feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("RSS feed returned status %d", resp.StatusCode)
	}

	return Parse(resp.Body)
}

// Parse reads RSS XML from r and returns the items.
func Parse(r io.Reader) ([]Item, error) {
	var feed rssFeed
	dec := xml.NewDecoder(r)
	if err := dec.Decode(&feed); err != nil {
		return nil, fmt.Errorf("parsing RSS XML: %w", err)
	}
	return feed.Channel.Items, nil
}

// RandomItem returns a random item from the fetched feed.
func RandomItem(items []Item) (Item, error) {
	if len(items) == 0 {
		return Item{}, fmt.Errorf("no items in feed")
	}
	return items[rand.Intn(len(items))], nil
}
