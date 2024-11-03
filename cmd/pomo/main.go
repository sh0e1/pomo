package main

import (
	"context"
	"os"

	"github.com/sh0e1/pomo/internal/cmd"
)

func main() {
	os.Exit(cmd.Run(context.Background()))
}
