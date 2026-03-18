package cmd

import (
	"strings"
	"testing"
)

func TestFormatCommitMessage(t *testing.T) {
	validTypes := map[string]bool{
		"feat": true, "fix": true, "chore": true, "refactor": true,
		"docs": true, "style": true, "perf": true, "test": true, "ci": true,
	}

	tests := []struct {
		name  string
		input string
	}{
		{"simple headline", "Area Man Passionate Defender Of What He Imagines Constitution To Be"},
		{"short headline", "No Way To Prevent This"},
		{"headline with punctuation", "Nation's Dog Owners Exposed As Tiniest Bit Racist"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatCommitMessage(tt.input)

			// Must contain a colon separator
			parts := strings.SplitN(got, ": ", 2)
			if len(parts) != 2 {
				t.Fatalf("formatCommitMessage(%q) = %q, want format 'type: message'", tt.input, got)
			}

			// Prefix must be a valid conventional commit type
			if !validTypes[parts[0]] {
				t.Errorf("formatCommitMessage(%q) prefix = %q, want one of %v", tt.input, parts[0], commitTypes)
			}

			// Message body must be lowercase version of input
			want := strings.ToLower(tt.input)
			if parts[1] != want {
				t.Errorf("formatCommitMessage(%q) body = %q, want %q", tt.input, parts[1], want)
			}
		})
	}
}

func TestFormatCommitMessageCoversAllTypes(t *testing.T) {
	seen := make(map[string]bool)
	// Run enough iterations to likely see all types
	for i := 0; i < 1000; i++ {
		msg := formatCommitMessage("test headline")
		parts := strings.SplitN(msg, ": ", 2)
		if len(parts) == 2 {
			seen[parts[0]] = true
		}
	}

	for _, ct := range commitTypes {
		if !seen[ct] {
			t.Errorf("commit type %q was never produced in 1000 iterations", ct)
		}
	}
}
