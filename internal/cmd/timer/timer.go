package timer

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/sh0e1/pomo/internal/elm/timer"
)

func NewCommand() *cobra.Command {
	var workInterval time.Duration

	cmd := &cobra.Command{
		Use:   "timer",
		Short: "start timer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return timer.Run(workInterval)
		},
	}

	cmd.Flags().DurationVarP(&workInterval, "work-interval", "w", 25*time.Minute, "work time interval")

	return cmd
}
