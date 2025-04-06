package app

import (
	"fmt"

	"github.com/justedd/hwgl-hw-1-1/internal/controller"
	"github.com/justedd/hwgl-hw-1-1/internal/usecase/counter"

	"log/slog"
)

type App struct {
	logger     *slog.Logger
	controller *controller.Controller
}

type Args struct {
	FileName string
	Top      uint
}

func New(logger *slog.Logger) (*App, error) {
	wordCounter, err := counter.New(logger)

	if err != nil {
		logger.Error("App New: initialization error", slog.Any("err", err))
		fmt.Println("Initialization error")

		return nil, fmt.Errorf("App New: %w", err)
	}

	return &App{
		logger:     logger,
		controller: controller.New(logger, wordCounter),
	}, nil
}

func (app *App) Run(rawArgs []string) {
	app.controller.HandleCall(rawArgs)
}
