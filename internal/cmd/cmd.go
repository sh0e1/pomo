package cmd

import (
	"context"
	"os"

	"github.com/sh0e1/pomo/internal/cmd/root"
)

type exitCode int

const (
	exitCodeOK    exitCode = 0
	exitCodeError exitCode = 1
)

func Run(ctx context.Context) exitCode {
	cmd := root.NewCommand()
	if err := cmd.ExecuteContext(ctx); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		return exitCodeError
	}
	return exitCodeOK
}
