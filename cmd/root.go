package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/maxbeizer/gh-onion/headlines"
	"github.com/maxbeizer/gh-onion/output"
	"github.com/spf13/cobra"
)

var (
	jsonFlag bool
	jqExpr   string
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gh-onion",
		Short: "Random Onion headlines in your terminal",
		Long:  "A gh CLI extension that brings The Onion to your terminal. Bundled dataset, zero API dependencies, pure fun.",
		RunE:  runRandom,
	}

	cmd.PersistentFlags().BoolVar(&jsonFlag, "json", false, "Output as JSON")
	cmd.PersistentFlags().StringVar(&jqExpr, "jq", "", "Filter JSON output (e.g. .text)")

	cmd.AddCommand(newMotdCmd())
	cmd.AddCommand(newSearchCmd())
	cmd.AddCommand(newFreshCmd())

	return cmd
}

func runRandom(cmd *cobra.Command, args []string) error {
	h, err := headlines.Random()
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

	fmt.Println(output.Box(h.Text, 60))
	fmt.Printf("\n  — %s\n", h.Source())
	return nil
}

// Execute runs the root command with signal handling.
func Execute() error {
	userMessages := log.New(os.Stderr, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		for sig := range c {
			userMessages.Printf("received signal %v", sig)
			cancel()
		}
	}()

	return newRootCmd().ExecuteContext(ctx)
}
