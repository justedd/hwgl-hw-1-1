package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/justedd/hwgl-hw-1-1/internal/usecase/counter"
)

type Args struct {
	FileName string
	Top      uint
}

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

	args, err := parseArgs(os.Args)

	if err != nil {
		logger.Error("main: args error", slog.Any("err", err))
		fmt.Println("Argument error")

		return
	}

	top, err := wordCounter.FileTop(args.Top, args.FileName)

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

func parseArgs(args []string) (*Args, error) {
	if len(args) != 3 {
		return nil, errors.New("wrong usage")
	}

	top, err := strconv.ParseUint(args[1], 10, 32)

	if err != nil {
		return nil, errors.New("invalid top number")
	}

	if len(args[2]) < 1 {
		return nil, errors.New("filename looks wrong")
	}

	return &Args{
		FileName: args[2],
		Top:      uint(top),
	}, nil
}
