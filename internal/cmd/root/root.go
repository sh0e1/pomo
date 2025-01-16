package root

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sh0e1/pomo/internal/cmd/timer"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Pomo!")
		},
	}

	cmd.AddCommand(timer.NewCommand())
	return cmd
}
