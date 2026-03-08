package output

import (
	"bytes"
	"strings"
	"testing"
)

func TestRenderJSON_object(t *testing.T) {
	var buf bytes.Buffer
	h := Headline{Text: "Test headline", IsOnion: true, Source: "stub"}
	if err := RenderJSON(&buf, h, ""); err != nil {
		t.Fatalf("RenderJSON() error = %v", err)
	}
	if !strings.Contains(buf.String(), `"text": "Test headline"`) {
		t.Errorf("expected JSON with text field, got %s", buf.String())
	}
}

func TestRenderJSON_jqFilter(t *testing.T) {
	var buf bytes.Buffer
	h := Headline{Text: "Test headline", IsOnion: true}
	if err := RenderJSON(&buf, h, ".text"); err != nil {
		t.Fatalf("RenderJSON() error = %v", err)
	}
	got := strings.TrimSpace(buf.String())
	if got != "Test headline" {
		t.Errorf("RenderJSON(.text) = %q, want %q", got, "Test headline")
	}
}

func TestRenderJSON_jqFilterArray(t *testing.T) {
	var buf bytes.Buffer
	items := []Headline{
		{Text: "One", IsOnion: true},
		{Text: "Two", IsOnion: false},
	}
	if err := RenderJSON(&buf, items, ".text"); err != nil {
		t.Fatalf("RenderJSON() error = %v", err)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 2 || lines[0] != "One" || lines[1] != "Two" {
		t.Errorf("expected [One, Two], got %v", lines)
	}
}

func TestBox(t *testing.T) {
	box := Box("Hello World", 30)
	if !strings.Contains(box, "╔") || !strings.Contains(box, "╚") {
		t.Error("Box() missing border characters")
	}
	if !strings.Contains(box, "Hello World") {
		t.Error("Box() missing content")
	}
}

func TestMOTD(t *testing.T) {
	got := MOTD("Test Headline")
	if !strings.Contains(got, "📰") {
		t.Error("MOTD() missing emoji")
	}
	if !strings.Contains(got, `"Test Headline"`) {
		t.Error("MOTD() missing quoted headline")
	}
}

func TestWrapText(t *testing.T) {
	tests := []struct {
		name  string
		text  string
		width int
		want  int // number of lines
	}{
		{"short", "hello", 60, 1},
		{"exact", "hello world", 11, 1},
		{"wraps", "this is a much longer sentence that should wrap", 20, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := wrapText(tt.text, tt.width)
			if len(lines) != tt.want {
				t.Errorf("wrapText(%q, %d) = %d lines, want %d", tt.text, tt.width, len(lines), tt.want)
			}
		})
	}
}
