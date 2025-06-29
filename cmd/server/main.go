package main

import (
	"command-service/internal/presentation"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		presentation.CommandDepend,
	).Run()
}
