package timer

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/sh0e1/pomo/internal/elm/timer"
)

func NewCommand() *cobra.Command {
	var (
		work       time.Duration
		shortBreak time.Duration
		longBreak  time.Duration
		rounds     int
	)

	cmd := &cobra.Command{
		Use:   "timer",
		Short: "start pomodoro timer 🍅",
		RunE: func(cmd *cobra.Command, args []string) error {
			return timer.Run(work)
		},
	}

	cmd.Flags().DurationVarP(&work, "work", "w", 25*time.Minute, "")
	cmd.Flags().DurationVarP(&shortBreak, "short-break", "s", 5*time.Minute, "")
	cmd.Flags().DurationVarP(&longBreak, "long-break", "l", 15*time.Minute, "")
	cmd.Flags().IntVarP(&rounds, "rounds", "r", 4, "")

	return cmd
}
