package app

import (
	"github.com/spf13/cobra"

	"github.com/ChrisRx/splits/cmd/splits/app/new"
	"github.com/ChrisRx/splits/cmd/splits/app/run"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "splits",
		Short: "",
	}

	cmd.AddCommand(
		new.NewCommand(),
		run.NewCommand(),
	)
	return cmd
}
