package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var service_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_rmCmd).Standalone()

	serviceCmd.AddCommand(service_rmCmd)

	carapace.Gen(service_rmCmd).PositionalAnyCompletion(action.ActionSerices())
}