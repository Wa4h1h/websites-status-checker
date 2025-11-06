package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New(root *cobra.Command, cmds ...*cobra.Command) *Cli {
	c := &Cli{
		rootCmd: root,
	}

	for _, cmd := range cmds {
		c.rootCmd.AddCommand(cmd)
	}

	return c
}

func (c *Cli) Run() error {
	if err := c.rootCmd.Execute(); err != nil {
		return fmt.Errorf("failed to start cli: %w", err)
	}

	return nil
}
