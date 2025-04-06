package controller

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/justedd/hwgl-hw-1-1/internal/usecase/counter"
)

type Controller struct {
	logger      *slog.Logger
	wordCounter *counter.Counter
}

type Args struct {
	FileName string
	Top      uint
}

var ErrWrongNumberOfArgs = errors.New("wrong number of args")
var ErrWrongFilename = errors.New("filename looks wrong")
var ErrWrongTop = errors.New("wrong top number")

func New(logger *slog.Logger, wordCounter *counter.Counter) *Controller {
	return &Controller{
		logger:      logger,
		wordCounter: wordCounter,
	}
}

func (c *Controller) HandleCall(rawArgs []string) {
	args, err := c.parseArgs(rawArgs)

	if err != nil {
		c.logger.Error("HandleCall: args error: ", slog.Any("err", err))
		fmt.Println("Argument error, usage: `topcounter 5 text.txt`")

		return
	}

	top, err := c.wordCounter.FileTop(args.Top, args.FileName)

	if err != nil {
		c.logger.Error("HandleCall: calculating top error", slog.Any("err", err))

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

func (c *Controller) parseArgs(args []string) (*Args, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("parseArgs: %w", ErrWrongNumberOfArgs)
	}

	top, err := strconv.ParseUint(args[1], 10, 32)

	if err != nil {
		return nil, fmt.Errorf("parseArgs: %w", ErrWrongTop)
	}

	if len(args[2]) < 1 {
		return nil, fmt.Errorf("parseArgs: %w", ErrWrongFilename)
	}

	return &Args{
		FileName: args[2],
		Top:      uint(top),
	}, nil
}
