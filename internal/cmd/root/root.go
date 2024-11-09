package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Pomo!")
		},
	}
	return cmd
}
