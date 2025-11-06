package main

import (
	"os"

	"github.com/Wa4h1h/websites-status-checker/internal/cli"
	"github.com/Wa4h1h/websites-status-checker/internal/cli/commands/check"
	"github.com/Wa4h1h/websites-status-checker/internal/statuschecker"
	"github.com/spf13/cobra"
)

func main() {
	c := cli.New(&cobra.Command{
		Use:   "wsc",
		Short: "A fast concurrent CLI tool to check multiple URLs’ HTTP status with configurable concurrency, timeouts, and detailed results summary.",
		Long:  "A fast concurrent CLI tool to check multiple URLs’ HTTP status with configurable concurrency, timeouts, and detailed results summary.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				_ = cmd.Usage()
				return
			}
		},
	}, check.New(statuschecker.New()))

	if err := c.Run(); err != nil {
		os.Stderr.Write([]byte(err.Error()))
	}
}
