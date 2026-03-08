package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/maxbeizer/gh-onion/headlines"
	"github.com/maxbeizer/gh-onion/output"
	"github.com/spf13/cobra"
)

func newSearchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "search <term>",
		Short: "Search headlines by keyword",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runSearch,
	}
}

func runSearch(cmd *cobra.Command, args []string) error {
	term := strings.Join(args, " ")
	results, err := headlines.Search(term)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		fmt.Fprintf(os.Stderr, "No headlines found matching %q\n", term)
		return nil
	}

	if jsonFlag || jqExpr != "" {
		items := make([]output.Headline, len(results))
		for i, h := range results {
			items[i] = output.Headline{Text: h.Text, IsOnion: h.IsOnion, Source: h.Source()}
		}
		return output.RenderJSON(os.Stdout, items, jqExpr)
	}

	fmt.Printf("Found %d headline(s) matching %q:\n\n", len(results), term)
	for _, h := range results {
		label := "🧅"
		if !h.IsOnion {
			label = "📰"
		}
		fmt.Printf("  %s %s\n", label, h.Text)
	}
	return nil
}
