package cmd

import (
	"fmt"
	"os"

	"github.com/maxbeizer/gh-onion/headlines"
	"github.com/maxbeizer/gh-onion/output"
	"github.com/spf13/cobra"
)

func newMotdCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "motd",
		Short: "Headline of the day for your shell RC",
		Long:  "Display a deterministic daily Onion headline, perfect for .zshrc or .bashrc.",
		RunE:  runMotd,
	}
}

func runMotd(cmd *cobra.Command, args []string) error {
	h, err := headlines.MOTD()
	if err != nil {
		return err
	}

	if jsonFlag || jqExpr != "" {
		return output.RenderJSON(os.Stdout, output.Headline{
			Text:    h.Text,
			IsOnion: h.IsOnion,
			Source:  h.Source(),
		}, jqExpr)
	}

	fmt.Println(output.MOTD(h.Text))
	return nil
}
