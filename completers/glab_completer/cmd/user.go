package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user <command> [flags]",
	Short: "Interact with user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userCmd).Standalone()
	rootCmd.AddCommand(userCmd)
}
