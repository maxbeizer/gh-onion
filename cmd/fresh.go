package cmd

import (
	"fmt"
	"os"

	"github.com/maxbeizer/gh-onion/output"
	"github.com/maxbeizer/gh-onion/rss"
	"github.com/spf13/cobra"
)

func newFreshCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "fresh",
		Short: "Fetch a fresh headline from The Onion's RSS feed",
		Long:  "Fetches The Onion's RSS feed and displays a random headline. Requires internet access.",
		RunE:  runFresh,
	}
}

func runFresh(cmd *cobra.Command, args []string) error {
	items, err := rss.Fetch()
	if err != nil {
		return fmt.Errorf("could not fetch fresh headlines (are you online?): %w", err)
	}

	item, err := rss.RandomItem(items)
	if err != nil {
		return err
	}

	if jsonFlag || jqExpr != "" {
		return output.RenderJSON(os.Stdout, output.Headline{
			Text:    item.Title,
			IsOnion: true,
			Source:  "The Onion (RSS)",
		}, jqExpr)
	}

	fmt.Println(output.Box(item.Title, 60))
	fmt.Printf("\n  🔗 %s\n", item.Link)
	return nil
}
