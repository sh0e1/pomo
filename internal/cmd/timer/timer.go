package timer

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/sh0e1/pomo/internal/elm/timer"
)

func NewCommand() *cobra.Command {
	cfg := &timer.Config{}

	cmd := &cobra.Command{
		Use:   "timer",
		Short: "start pomodoro timer 🍅",
		RunE: func(cmd *cobra.Command, args []string) error {
			return timer.Run(cfg)
		},
	}

	cmd.Flags().DurationVarP(&cfg.WorkInterval, "work", "w", 25*time.Minute, "")
	cmd.Flags().DurationVarP(&cfg.ShortBreakInterval, "short-break", "s", 5*time.Minute, "")
	cmd.Flags().DurationVarP(&cfg.LongBreakInterval, "long-break", "l", 15*time.Minute, "")
	cmd.Flags().IntVarP(&cfg.Rounds, "rounds", "r", 4, "")

	return cmd
}
