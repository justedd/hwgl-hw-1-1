package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

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

	top, err := wordCounter.FileTop(5, "text.txt")

	if err != nil {
		logger.Error("main: calculating top error", slog.Any("err", err))

		if errors.Is(err, counter.ErrFileOpen) {
			fmt.Println("Error while openning file")
		} else {
			fmt.Println("Internal error")
		}
		return
	}

	for i := range top {
		line := fmt.Sprintf("%s: %d", top[i].Word, top[i].Count)
		fmt.Println(line)
	}
}
