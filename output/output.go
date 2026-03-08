package output

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Headline is the minimal interface for output rendering.
type Headline struct {
	Text    string `json:"text"`
	IsOnion bool   `json:"is_onion"`
	Source  string `json:"source,omitempty"`
}

// RenderJSON writes data as JSON to w. If jqExpr is non-empty, it filters to
// matching top-level keys (lightweight; not full jq).
func RenderJSON(w io.Writer, data any, jqExpr string) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	if jqExpr != "" {
		return renderJQFilter(w, b, jqExpr)
	}
	_, err = fmt.Fprintln(w, string(b))
	return err
}

// renderJQFilter applies a simple top-level field filter like ".text" or ".is_onion".
func renderJQFilter(w io.Writer, raw []byte, expr string) error {
	expr = strings.TrimPrefix(expr, ".")

	// Try as single object first
	var obj map[string]any
	if err := json.Unmarshal(raw, &obj); err == nil {
		if val, ok := obj[expr]; ok {
			return writeJSONValue(w, val)
		}
		return fmt.Errorf("field %q not found", expr)
	}

	// Try as array of objects
	var arr []map[string]any
	if err := json.Unmarshal(raw, &arr); err == nil {
		for _, item := range arr {
			if val, ok := item[expr]; ok {
				if err := writeJSONValue(w, val); err != nil {
					return err
				}
			}
		}
		return nil
	}

	return fmt.Errorf("unsupported jq expression: %s", expr)
}

func writeJSONValue(w io.Writer, val any) error {
	switch v := val.(type) {
	case string:
		_, err := fmt.Fprintln(w, v)
		return err
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(w, string(b))
		return err
	}
}

// Box wraps text in a Unicode box-drawing frame.
func Box(text string, width int) string {
	if width < 4 {
		width = 60
	}
	innerWidth := width - 4 // 2 for border + 2 for padding

	lines := wrapText(text, innerWidth)

	var sb strings.Builder
	sb.WriteString("  ╔" + strings.Repeat("═", innerWidth+2) + "╗\n")
	for _, line := range lines {
		padding := innerWidth - runeLen(line)
		if padding < 0 {
			padding = 0
		}
		sb.WriteString("  ║ " + line + strings.Repeat(" ", padding) + " ║\n")
	}
	sb.WriteString("  ╚" + strings.Repeat("═", innerWidth+2) + "╝")
	return sb.String()
}

// MOTD formats a headline as a "message of the day".
func MOTD(text string) string {
	return fmt.Sprintf("📰 Today's Headline from America's Finest News Source:\n   \"%s\"", text)
}

func wrapText(text string, width int) []string {
	if width <= 0 {
		return []string{text}
	}
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{""}
	}

	var lines []string
	current := words[0]
	for _, word := range words[1:] {
		if runeLen(current)+1+runeLen(word) <= width {
			current += " " + word
		} else {
			lines = append(lines, current)
			current = word
		}
	}
	lines = append(lines, current)
	return lines
}

func runeLen(s string) int {
	return len([]rune(s))
}
