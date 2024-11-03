package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Pomo!")
		},
	}
}
