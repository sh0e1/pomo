package main

import (
	"context"
	"os"

	"github.com/sh0e1/pomo/internal/cmd"
)

func main() {
	code := cmd.Run(context.Background())
	os.Exit(int(code))
}
