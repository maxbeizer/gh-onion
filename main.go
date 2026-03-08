package main

import (
	"os"

	"github.com/maxbeizer/gh-onion/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
