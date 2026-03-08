package rss

import (
	"strings"
	"testing"
)

const testFeed = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <title>The Onion</title>
    <item>
      <title>Area Man Discovers New Emotion</title>
      <link>https://www.theonion.com/area-man-discovers-new-emotion</link>
      <pubDate>Mon, 01 Jan 2024 00:00:00 GMT</pubDate>
    </item>
    <item>
      <title>Nation's Dogs Vow To Continue Barking At Nothing</title>
      <link>https://www.theonion.com/nations-dogs-barking</link>
      <pubDate>Tue, 02 Jan 2024 00:00:00 GMT</pubDate>
    </item>
    <item>
      <title>Study Finds Sitting Up Warming Up</title>
      <link>https://www.theonion.com/study-sitting</link>
      <pubDate>Wed, 03 Jan 2024 00:00:00 GMT</pubDate>
    </item>
  </channel>
</rss>`

func TestParse(t *testing.T) {
	items, err := Parse(strings.NewReader(testFeed))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if len(items) != 3 {
		t.Fatalf("Parse() returned %d items, want 3", len(items))
	}
	if items[0].Title != "Area Man Discovers New Emotion" {
		t.Errorf("items[0].Title = %q, want %q", items[0].Title, "Area Man Discovers New Emotion")
	}
	if items[0].Link == "" {
		t.Error("items[0].Link is empty")
	}
}

func TestParse_empty(t *testing.T) {
	empty := `<?xml version="1.0"?><rss><channel></channel></rss>`
	items, err := Parse(strings.NewReader(empty))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if len(items) != 0 {
		t.Errorf("Parse() returned %d items, want 0", len(items))
	}
}

func TestParse_invalid(t *testing.T) {
	_, err := Parse(strings.NewReader("not xml"))
	if err == nil {
		t.Error("Parse() expected error for invalid XML")
	}
}

func TestRandomItem(t *testing.T) {
	items, _ := Parse(strings.NewReader(testFeed))
	item, err := RandomItem(items)
	if err != nil {
		t.Fatalf("RandomItem() error = %v", err)
	}
	if item.Title == "" {
		t.Error("RandomItem() returned empty title")
	}
}

func TestRandomItem_empty(t *testing.T) {
	_, err := RandomItem(nil)
	if err == nil {
		t.Error("RandomItem(nil) expected error")
	}
}
