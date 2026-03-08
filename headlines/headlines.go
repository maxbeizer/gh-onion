package headlines

import (
	"encoding/csv"
	"embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:embed data/onion.csv
var dataFS embed.FS

// Headline represents a single headline entry.
type Headline struct {
	Text    string `json:"text"`
	IsOnion bool   `json:"is_onion"`
}

// Source returns "The Onion" or "Real News" based on the label.
func (h Headline) Source() string {
	if h.IsOnion {
		return "The Onion"
	}
	return "Real (not The Onion)"
}

var loaded []Headline

// All returns all headlines from the embedded dataset.
func All() ([]Headline, error) {
	if loaded != nil {
		return loaded, nil
	}

	f, err := dataFS.Open("data/onion.csv")
	if err != nil {
		return nil, fmt.Errorf("opening embedded CSV: %w", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parsing CSV: %w", err)
	}

	for i, rec := range records {
		if i == 0 {
			continue // skip header
		}
		if len(rec) < 2 {
			continue
		}
		isOnion := strings.TrimSpace(rec[1]) == "1"
		loaded = append(loaded, Headline{
			Text:    strings.TrimSpace(rec[0]),
			IsOnion: isOnion,
		})
	}

	return loaded, nil
}

// Random returns a random headline.
func Random() (Headline, error) {
	all, err := All()
	if err != nil {
		return Headline{}, err
	}
	if len(all) == 0 {
		return Headline{}, fmt.Errorf("no headlines loaded")
	}
	return all[rand.Intn(len(all))], nil
}

// MOTD returns a deterministic headline for the current day.
func MOTD() (Headline, error) {
	all, err := All()
	if err != nil {
		return Headline{}, err
	}
	if len(all) == 0 {
		return Headline{}, fmt.Errorf("no headlines loaded")
	}

	now := time.Now()
	dayIndex := (now.Year()*1000 + now.YearDay()) % len(all)
	return all[dayIndex], nil
}

// Search returns headlines matching the term (case-insensitive substring).
func Search(term string) ([]Headline, error) {
	all, err := All()
	if err != nil {
		return nil, err
	}

	lower := strings.ToLower(term)
	var results []Headline
	for _, h := range all {
		if strings.Contains(strings.ToLower(h.Text), lower) {
			results = append(results, h)
		}
	}
	return results, nil
}
