package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/justedd/hwgl-hw-1-1/internal/app"
	"github.com/justedd/hwgl-hw-1-1/internal/usecase/counter"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	wordCounter, err := counter.NewCounter(logger)

	if err != nil {
		logger.Error("main: initialization error", slog.Any("err", err))
		fmt.Println("Initialization error")

		return
	}

	cli := app.New(logger, wordCounter)
	cli.Run(os.Args)
}
