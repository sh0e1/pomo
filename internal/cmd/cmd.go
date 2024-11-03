package cmd

import (
	"context"
	"fmt"
)

func Run(ctx context.Context) int {
	fmt.Println("Hello, Pomo!")
	return 0
}
