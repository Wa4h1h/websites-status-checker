package cli

import "github.com/spf13/cobra"

type Cli struct {
	rootCmd     *cobra.Command
	subCommands []*cobra.Command
}
