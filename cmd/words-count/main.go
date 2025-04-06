package main

import (
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

	counter, err := counter.NewCounter(logger)
	if err != nil {
		fmt.Printf("Initialization error: %v", err)
		return
	}

	top, err := counter.FileTop(5, "text.txt")

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	for i := range top {
		line := fmt.Sprintf("%s: %d", top[i].Word, top[i].Count)
		fmt.Println(line)
	}
}
