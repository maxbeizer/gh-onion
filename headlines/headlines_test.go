package headlines

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	// Reset cache for clean test
	loaded = nil

	all, err := All()
	if err != nil {
		t.Fatalf("All() error = %v", err)
	}
	if len(all) == 0 {
		t.Fatal("All() returned empty slice")
	}

	// Check we have both Onion and real headlines
	var onion, real int
	for _, h := range all {
		if h.IsOnion {
			onion++
		} else {
			real++
		}
	}
	if onion == 0 {
		t.Error("expected at least one Onion headline")
	}
	if real == 0 {
		t.Error("expected at least one real headline")
	}
}

func TestRandom(t *testing.T) {
	loaded = nil
	h, err := Random()
	if err != nil {
		t.Fatalf("Random() error = %v", err)
	}
	if h.Text == "" {
		t.Error("Random() returned empty headline")
	}
}

func TestMOTD(t *testing.T) {
	loaded = nil
	h, err := MOTD()
	if err != nil {
		t.Fatalf("MOTD() error = %v", err)
	}
	if h.Text == "" {
		t.Error("MOTD() returned empty headline")
	}

	// MOTD should be deterministic for same day
	h2, _ := MOTD()
	if h.Text != h2.Text {
		t.Error("MOTD() not deterministic within same day")
	}
}

func TestSearch(t *testing.T) {
	loaded = nil
	tests := []struct {
		name      string
		term      string
		wantMin   int
		wantEmpty bool
	}{
		{"finds man", "man", 1, false},
		{"case insensitive", "MAN", 1, false},
		{"no results", "xyzzyzzyx", 0, true},
		{"finds florida", "florida", 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := Search(tt.term)
			if err != nil {
				t.Fatalf("Search(%q) error = %v", tt.term, err)
			}
			if tt.wantEmpty && len(results) != 0 {
				t.Errorf("Search(%q) = %d results, want 0", tt.term, len(results))
			}
			if !tt.wantEmpty && len(results) < tt.wantMin {
				t.Errorf("Search(%q) = %d results, want >= %d", tt.term, len(results), tt.wantMin)
			}
			// Verify all results contain the search term
			for _, r := range results {
				if !strings.Contains(strings.ToLower(r.Text), strings.ToLower(tt.term)) {
					t.Errorf("result %q doesn't contain term %q", r.Text, tt.term)
				}
			}
		})
	}
}

func TestHeadline_Source(t *testing.T) {
	tests := []struct {
		isOnion bool
		want    string
	}{
		{true, "The Onion"},
		{false, "Real (not The Onion)"},
	}
	for _, tt := range tests {
		h := Headline{Text: "test", IsOnion: tt.isOnion}
		if got := h.Source(); got != tt.want {
			t.Errorf("Source() = %q, want %q", got, tt.want)
		}
	}
}
