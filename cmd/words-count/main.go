package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/justedd/hwgl-hw-1-1/internal/app"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)

	cli, err := app.New(logger)

	if err != nil {
		logger.Error("main: initialization error", slog.Any("err", err))
		fmt.Println("Initialization error")

		return
	}

	cli.Run(os.Args)
}
