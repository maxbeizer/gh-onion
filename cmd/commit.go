package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/maxbeizer/gh-onion/headlines"
	"github.com/maxbeizer/gh-onion/output"
	"github.com/spf13/cobra"
)

// commitTypes are conventional-commit prefixes for maximum chaos.
var commitTypes = []string{
	"feat", "fix", "chore", "refactor", "docs",
	"style", "perf", "test", "ci",
}

func newCommitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "commit",
		Short: "Output a headline as a chaotic commit message",
		Long:  "Picks a random Onion headline and formats it as a conventional commit message.\nPipe to git commit -m for the brave.",
		RunE:  runCommit,
	}
}

// formatCommitMessage formats a headline as a conventional commit message.
func formatCommitMessage(text string) string {
	prefix := commitTypes[rand.Intn(len(commitTypes))]
	return fmt.Sprintf("%s: %s", prefix, strings.ToLower(text))
}

func runCommit(cmd *cobra.Command, args []string) error {
	h, err := headlines.Random()
	if err != nil {
		return err
	}

	msg := formatCommitMessage(h.Text)

	if jsonFlag || jqExpr != "" {
		return output.RenderJSON(os.Stdout, output.Headline{
			Text:    msg,
			IsOnion: h.IsOnion,
			Source:  h.Source(),
		}, jqExpr)
	}

	fmt.Println(msg)
	return nil
}
